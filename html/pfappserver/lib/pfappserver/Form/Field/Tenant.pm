package pfappserver::Form::Field::Tenant;

=head1 NAME

pfappserver::Form::Field::Tenant -

=head1 DESCRIPTION

pfappserver::Form::Field::Tenant

=cut

use strict;
use warnings;
use HTML::FormHandler::Moose;
extends 'HTML::FormHandler::Field::Select';
use pf::dal::tenant;
use pf::constants qw($DEFAULT_TENANT_ID);
has 'no_global' => (
    is => 'rw',
    default => 0,
);

has '+default' => (
    default => $DEFAULT_TENANT_ID,
);

sub build_options {
    my $self = shift;
    my @tenants = map { (!$self->no_global || $_->{id} != 0) ? ( { value => $_->{id}, label => $_->{name} }) : () } @{pf::dal::tenant->search->all};
    return \@tenants;
}


=head1 AUTHOR

Inverse inc. <info@inverse.ca>

=head1 COPYRIGHT

Copyright (C) 2005-2020 Inverse inc.

=head1 LICENSE

This program is free software; you can redistribute it and/or
modify it under the terms of the GNU General Public License
as published by the Free Software Foundation; either version 2
of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program; if not, write to the Free Software
Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301,
USA.

=cut

1;
