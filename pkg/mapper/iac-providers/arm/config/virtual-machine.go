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
	networkInterfaces = "networkInterfaces"
)

// virtualMachineConfig holds config for azurerm_virtual_machine
type virtualMachineConfig struct {
	config

	NetworkInterfaceIDs []string `json:"network_interface_ids"`
}

// VirtualMachineConfig returns config for azurerm_virtual_machine
func VirtualMachineConfig(r types.Resource, params map[string]interface{}) interface{} {
	cf := virtualMachineConfig{
		config: config{
			Location: fn.LookUp(nil, params, r.Location).(string),
			Name:     fn.LookUp(nil, params, r.Name).(string),
			Tags:     r.Tags,
		},
		NetworkInterfaceIDs: []string{},
	}

	profile := convert.ToMap(r.Properties, networkProfile)
	if interfaces, ok := profile[networkInterfaces].([]interface{}); ok {
		for _, fs := range interfaces {
			iFace := fs.(map[string]interface{})
			cf.NetworkInterfaceIDs = append(cf.NetworkInterfaceIDs, iFace["id"].(string))
		}
	}
	return cf
}
