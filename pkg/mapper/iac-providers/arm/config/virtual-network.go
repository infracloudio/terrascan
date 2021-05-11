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
	subnets              = "subnets"
	properties           = "properties"
	addressPrefix        = "addressPrefix"
	networkSecurityGroup = "networkSecurityGroup"
)

// virtualNetworkConfig holds config for azurerm_virtual_network
type virtualNetworkConfig struct {
	config

	Subnets []subnet `json:"subnet"`
}

type subnet struct {
	Name          string `json:"name"`
	AddressPrefix string `json:"address_prefix"`
	SecurityGroup string `json:"security_group,omitempty"`
}

// VirtualNetworkConfig returns config for azurerm_virtual_network
func VirtualNetworkConfig(r types.Resource, vars, params map[string]interface{}) interface{} {
	cf := virtualNetworkConfig{
		config: config{
			Location: fn.LookUp(vars, params, r.Location).(string),
			Name:     fn.LookUp(vars, params, r.Name).(string),
			Tags:     r.Tags,
		},
	}

	subnets := convert.ToSlice(r.Properties, subnets)
	for _, ss := range subnets {
		s := ss.(map[string]interface{})
		prop := convert.ToMap(s, properties)

		sub := subnet{
			Name:          fn.LookUp(vars, params, s[name].(string)).(string),
			AddressPrefix: fn.LookUp(vars, params, prop[addressPrefix].(string)).(string),
		}

		if nsg := convert.ToMap(prop, networkSecurityGroup); nsg != nil {
			if sg, ok := fn.LookUp(vars, params, nsg["id"].(string)).(string); ok {
				sub.SecurityGroup = sg
			}
		}
		cf.Subnets = append(cf.Subnets, sub)
	}
	return cf
}
