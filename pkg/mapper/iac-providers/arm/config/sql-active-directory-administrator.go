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
	login    = "login"
	sid      = "sid"
	tenantID = "tenantId"
)

// sqlActiveDirectoryAdministratorConfig holds config for azurerm_sql_active_directory_administrator
type sqlActiveDirectoryAdministratorConfig struct {
	config

	Login    string `json:"login"`
	TenantID string `json:"tenant_id"`
	ObjectID string `json:"object_id"`
}

// SQLActiveDirectoryAdministratorConfig returns config for azurerm_sql_active_directory_administrator
func SQLActiveDirectoryAdministratorConfig(r types.Resource, vars, params map[string]interface{}) interface{} {
	return sqlActiveDirectoryAdministratorConfig{
		config: config{
			Location: fn.LookUp(nil, params, r.Location).(string),
			Name:     fn.LookUp(nil, params, r.Name).(string),
		},
		Login:    fn.LookUp(vars, params, convert.ToString(r.Properties, login)).(string),
		ObjectID: fn.LookUp(vars, params, convert.ToString(r.Properties, sid)).(string),
		TenantID: fn.LookUp(vars, params, convert.ToString(r.Properties, tenantID)).(string),
	}
}
