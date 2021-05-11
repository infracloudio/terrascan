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

const enableSoftDelete = "enableSoftDelete"

// keyVaultConfig holds config for azurerm_key_vault
type keyVaultConfig struct {
	config

	TenantID          string `json:"tenant_id"`
	SoftDeleteEnabled bool   `json:"soft_delete_enabled"`
}

// KeyVaultConfig returns config for azurerm_key_vault
func KeyVaultConfig(r types.Resource, params map[string]interface{}) interface{} {
	return keyVaultConfig{
		config: config{
			Location: fn.LookUp(nil, params, r.Location).(string),
			Name:     fn.LookUp(nil, params, r.Name).(string),
		},
		SoftDeleteEnabled: convert.ToBool(r.Properties, enableSoftDelete),
		TenantID:          convert.ToString(params, tenantID),
	}
}
