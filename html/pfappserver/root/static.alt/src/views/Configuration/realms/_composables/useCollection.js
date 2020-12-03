import { computed, toRefs } from '@vue/composition-api'
import i18n from '@/utils/locale'
import {
  defaultsFromMeta as useItemDefaults
} from '../../_config/'

const useItemTitle = (props) => {
  const {
    id,
    isClone,
    isNew
  } = toRefs(props)
  return computed(() => {
    switch (true) {
      case !isNew.value && !isClone.value:
        return i18n.t('Realm: <code>{id}</code>', { id: id.value })
      case isClone.value:
        return i18n.t('Clone Realm: <code>{id}</code>', { id: id.value })
      default:
        return i18n.t('New Realm')
    }
  })
}

const useRouter = (props, context, form) => {
  const {
    id
  } = toRefs(props)
  const { root: { $router } = {} } = context
  return {
    goToCollection: () => $router.push({ name: 'realms' }),
    goToItem: () => $router.push({ name: 'realm', params: { id: form.value.id || id.value } }),
    goToClone: () => $router.push({ name: 'cloneRealm', params: { id: id.value } }),
  }
}

const useStore = (props, context, form) => {
  const {
    id,
    isClone
  } = toRefs(props)
  const { root: { $store } = {} } = context
  return {
    isLoading: computed(() => $store.getters['$_realms/isLoading']),
    getOptions: () => $store.dispatch('$_realms/options'),
    createItem: () => $store.dispatch('$_realms/createRealm', form.value),
    deleteItem: () => $store.dispatch('$_realms/deleteRealm', id.value),
    getItem: () => $store.dispatch('$_realms/getRealm', id.value).then(item => {
      if (isClone.value) {
        item.id = `${item.id}-${i18n.t('copy')}`
        item.not_deletable = false
      }
      return item
    }),
    updateItem: () => $store.dispatch('$_realms/updateRealm', form.value),
  }
}

export default {
  useItemDefaults,
  useItemTitle,
  useRouter,
  useStore,
}