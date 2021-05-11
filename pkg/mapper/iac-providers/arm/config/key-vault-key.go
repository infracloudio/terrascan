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
	"time"

	"github.com/accurics/terrascan/pkg/mapper/convert"
	fn "github.com/accurics/terrascan/pkg/mapper/iac-providers/arm/functions"
	"github.com/accurics/terrascan/pkg/mapper/iac-providers/arm/types"
)

// keyVaultKeyConfig holds config for azurerm_key_vault_key
type keyVaultKeyConfig struct {
	config
	ExpirationDate string `json:"expiration_date"`
}

// KeyVaultKeyConfig returns config for azurerm_key_vault_key
func KeyVaultKeyConfig(r types.Resource, params map[string]interface{}) interface{} {
	cf := keyVaultKeyConfig{
		config: config{
			Location: fn.LookUp(nil, params, r.Location).(string),
			Name:     fn.LookUp(nil, params, r.Name).(string),
		},
	}

	attr := convert.ToMap(r.Properties, attributes)
	if i := attr["exp"]; i != nil {
		t := time.Unix(int64(i.(float64)), 0)
		cf.ExpirationDate = t.Format(time.RFC3339)
	}
	return cf
}
