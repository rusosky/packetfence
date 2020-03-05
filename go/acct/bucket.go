package main

import (
	"github.com/inverse-inc/packetfence/go/mac"
	"sync"
	"time"
)

type Session [16]byte

type BucketKey struct {
	TenantId int
	Mac      mac.Mac
}

type TimeBucketKey struct {
	SessionID  Session
	TimeBucket int64
}

type BandwidthBucket struct {
	InBytes  uint64
	OutBytes uint64
}

type BandwidthBuckets struct {
	lock    sync.RWMutex
	Buckets map[TimeBucketKey]BandwidthBucket
}

type Buckets struct {
	lock             sync.RWMutex
	BandwidthBuckets map[BucketKey]*BandwidthBuckets
}

func (b *Buckets) Add(tenantId int, mac mac.Mac, sessionID Session, timeBucket time.Time, in, out uint64) {
	bb := b.getOrAdd(tenantId, mac, sessionID, timeBucket)
	bb.Add(sessionID, timeBucket, in, out)
}

func (b *Buckets) Update(tenantId int, mac mac.Mac, sessionID Session, timeBucket time.Time, in, out uint64) {
	bb := b.getOrAdd(tenantId, mac, sessionID, timeBucket)
	bb.Update(sessionID, timeBucket, in, out)
}

func (b *Buckets) getOrAdd(tenantId int, mac mac.Mac, sessionID Session, timeBucket time.Time) *BandwidthBuckets {
	var bb *BandwidthBuckets
	var found bool
	key := BucketKey{TenantId: tenantId, Mac: mac}
	b.lock.RLock() //First try to get a read Only lock
	if bb, found = b.BandwidthBuckets[key]; !found {
		// Not Found drop read lock get write lock
		b.lock.RUnlock()
		b.lock.Lock()
		// Recheck if exist if not found add
		if bb, found = b.BandwidthBuckets[key]; !found {
			bb = &BandwidthBuckets{Buckets: make(map[TimeBucketKey]BandwidthBucket)}
			b.BandwidthBuckets[key] = bb
		}
		b.lock.Unlock()
	} else {
		b.lock.RUnlock()
	}
    return bb
}

func (b *Buckets) GetBucket(tenantId int, mac mac.Mac, sessionID Session, timeBucket time.Time) (BandwidthBucket, bool) {
	var bb *BandwidthBuckets
	var found bool
	key := BucketKey{TenantId: tenantId, Mac: mac}
	b.lock.RLock() //First try to get a read Only lock
	if bb, found = b.BandwidthBuckets[key]; !found {
		b.lock.RUnlock()
        return BandwidthBucket{}, false
    }
    b.lock.RUnlock()
    return bb.GetBucket(sessionID, timeBucket)
}

func (b *BandwidthBuckets) GetBucket(sessionID Session, timeBucket time.Time) (BandwidthBucket, bool) {
        key := TimeBucketKey{SessionID: sessionID, TimeBucket: timeBucket.UnixNano()}
        b.lock.RLock()
        defer b.lock.RUnlock()
        bk, found := b.Buckets[key]
        return bk, found
}

func (b *BandwidthBuckets) Add(sessionID Session, timeBucket time.Time, in, out uint64) {
	key := TimeBucketKey{SessionID: sessionID, TimeBucket: timeBucket.UnixNano()}
	b.lock.Lock()
	bb := b.Buckets[key]
	bb.InBytes += in
	bb.OutBytes += out
	b.Buckets[key] = bb
	b.lock.Unlock()
}

func (b *BandwidthBuckets) Update(sessionID Session, timeBucket time.Time, in, out uint64) {
	key := TimeBucketKey{SessionID: sessionID, TimeBucket: timeBucket.UnixNano()}
	bb := BandwidthBucket{InBytes: in, OutBytes: out}
	b.lock.Lock()
	for k, v := range b.Buckets {
		if k.SessionID != key.SessionID || k.TimeBucket == key.TimeBucket {
			continue
		}
		bb.InBytes -= v.InBytes
		bb.OutBytes -= v.OutBytes
	}
	b.Buckets[key] = bb
	b.lock.Unlock()
}

func NewBuckets() *Buckets {
	return &Buckets{BandwidthBuckets: make(map[BucketKey]*BandwidthBuckets)}
}
