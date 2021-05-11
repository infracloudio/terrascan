/*
    Copyright (C) 2020 Accurics, Inc.

	Licensed under the Apache License, Version 2.0 (the "License");
    you may not use this file except in compliance with the License.
    You may obtain a copy of the License at

		http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License.
*/

package config

import (
	"github.com/accurics/terrascan/pkg/mapper/convert"
	fn "github.com/accurics/terrascan/pkg/mapper/iac-providers/arm/functions"
	"github.com/accurics/terrascan/pkg/mapper/iac-providers/arm/types"
)

const (
	access               = "access"
	direction            = "direction"
	protocol             = "protocol"
	sourceAddressPrefix  = "sourceAddressPrefix"
	sourcePortRange      = "sourcePortRange"
	destinationPortRange = "destinationPortRange"
)

// networkSecurityRuleConfig holds config for azurerm_network_security_rule
type networkSecurityRuleConfig struct {
	config

	Access               string `json:"access"`
	Direction            bool   `json:"direction"`
	Protocol             string `json:"protocol"`
	SourceAddressPrefix  string `json:"source_address_prefix"`
	SourcePortRange      string `json:"source_port_range"`
	DestinationPortRange string `json:"destination_port_range"`
}

// NetworkSecurityRuleConfig returns config for azurerm_network_security_rule
func NetworkSecurityRuleConfig(r types.Resource, params map[string]interface{}) interface{} {
	return networkSecurityRuleConfig{
		config: config{
			Location: fn.LookUp(nil, params, r.Location).(string),
			Name:     fn.LookUp(nil, params, r.Name).(string),
		},
		Access:               convert.ToString(r.Properties, access),
		Direction:            convert.ToBool(r.Properties, direction),
		Protocol:             convert.ToString(r.Properties, protocol),
		SourceAddressPrefix:  convert.ToString(r.Properties, sourceAddressPrefix),
		SourcePortRange:      convert.ToString(r.Properties, sourcePortRange),
		DestinationPortRange: convert.ToString(r.Properties, destinationPortRange),
	}
}
