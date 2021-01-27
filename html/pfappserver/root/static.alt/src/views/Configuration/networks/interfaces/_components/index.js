import { BaseViewCollectionItem } from '../../../_components/new/'
import {
  BaseFormButtonBar,
  BaseFormGroupChosenMultiple,
  BaseFormGroupChosenOne,
  BaseFormGroupInput,
  BaseFormGroupInputNumber,
  BaseFormGroupToggleDisabledEnabled
} from '@/components/new/'
import TheForm from './TheForm'
import TheView from './TheView'

export {
  BaseViewCollectionItem              as BaseView,
  BaseFormButtonBar                   as FormButtonBar,

  BaseFormGroupChosenMultiple         as FormGroupAdditionalListeneningDaemons,
  BaseFormGroupToggleDisabledEnabled  as FormGroupCoa,
  BaseFormGroupInput                  as FormGroupDescription,
  BaseFormGroupToggleDisabledEnabled  as FormGroupDhcpdEnabled,
  BaseFormGroupInput                  as FormGroupDns,
  BaseFormGroupToggleDisabledEnabled  as FormGroupHighAvailability,
  BaseFormGroupInput                  as FormGroupIdentifier,
  BaseFormGroupInput                  as FormGroupIpAddress,
  BaseFormGroupInput                  as FormGroupIpv6Address,
  BaseFormGroupInputNumber            as FormGroupIpv6Prefix,
  BaseFormGroupToggleDisabledEnabled  as FormGroupNatEnabled,
  BaseFormGroupToggleDisabledEnabled  as FormGroupNetflowAccountingEnabled,
  BaseFormGroupInput                  as FormGroupNetmask,
  BaseFormGroupInput                  as FormGroupRegNetwork,
  BaseFormGroupToggleDisabledEnabled  as FormGroupSplitNetwork,
  BaseFormGroupChosenOne              as FormGroupType,
  BaseFormGroupInput                  as FormGroupVlan,

  TheForm,
  TheView
}
