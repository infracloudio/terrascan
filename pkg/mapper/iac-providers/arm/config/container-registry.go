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

const adminUserEnabled = "adminUserEnabled"

// containerRegistryConfig holds config for azurerm_container_registry
type containerRegistryConfig struct {
	config

	SKU          string `json:"sku"`
	AdminEnabled bool   `json:"admin_enabled"`
}

// ContainerRegistryConfig returns config for azurerm_container_registry
func ContainerRegistryConfig(r types.Resource, params map[string]interface{}) interface{} {
	return containerRegistryConfig{
		config: config{
			Location: fn.LookUp(nil, params, r.Location).(string),
			Name:     fn.LookUp(nil, params, r.Name).(string),
		},
		SKU:          fn.LookUp(nil, params, r.SKU.Name).(string),
		AdminEnabled: fn.LookUp(nil, params, convert.ToString(r.Properties, adminUserEnabled)).(bool),
	}
}
