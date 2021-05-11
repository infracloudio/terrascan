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
	"strings"

	"github.com/accurics/terrascan/pkg/mapper/convert"
	fn "github.com/accurics/terrascan/pkg/mapper/iac-providers/arm/functions"
	"github.com/accurics/terrascan/pkg/mapper/iac-providers/arm/types"
)

const (
	administratorLogin         = "administratorLogin"
	administratorLoginPassword = "administratorLoginPassword"
	minimumTLSVersion          = "minimalTlsVersion"
	auditingSettings           = "auditingSettings"
	storageEndpoint            = "storageEndpoint"
	storageAccountAccessKey    = "storageAccountAccessKey"
	isStorageSecondaryKeyInUse = "isStorageSecondaryKeyInUse"
	retentionDays              = "retentionDays"
)

// managementLockConfig holds config for azurerm_mssql_server
type mssqlServerConfig struct {
	config

	AdministratorLogin         string `json:"administrator_login"`
	AdministratorLoginPassword string `json:"administrator_login_password"`
	MinimumTLSVersion          string `json:"minimum_tls_version"`

	ExtendedAuditingPolicy struct {
		StorageEndpoint                    string  `json:"storage_endpoint,omitempty"`
		StorageAccountAccessKey            string  `json:"storage_account_access_key,omitempty"`
		StorageAccountAccessKeyIsSecondary bool    `json:"storage_account_access_key_is_secondary,omitempty"`
		RetentionInDays                    float64 `json:"retention_in_days,omitempty"`
	} `json:"extended_auditing_policy,omitempty"`
}

// MSSQLServerConfig returns config for azurerm_mssql_server
func MSSQLServerConfig(r types.Resource, vars, params map[string]interface{}) interface{} {
	cf := mssqlServerConfig{
		config: config{
			Location: fn.LookUp(nil, params, r.Location).(string),
			Name:     fn.LookUp(nil, params, r.Name).(string),
		},
		AdministratorLogin:         fn.LookUp(vars, params, convert.ToString(r.Properties, administratorLogin)).(string),
		AdministratorLoginPassword: fn.LookUp(vars, params, convert.ToString(r.Properties, administratorLoginPassword)).(string),
		MinimumTLSVersion:          fn.LookUp(vars, params, convert.ToString(r.Properties, minimumTLSVersion)).(string),
	}

	for _, res := range r.Resources {
		if strings.EqualFold(res.Type, auditingSettings) {
			cf.ExtendedAuditingPolicy.StorageEndpoint = convert.ToString(res.Properties, storageEndpoint)
			cf.ExtendedAuditingPolicy.StorageAccountAccessKey = convert.ToString(res.Properties, storageAccountAccessKey)
			cf.ExtendedAuditingPolicy.StorageAccountAccessKeyIsSecondary = convert.ToBool(res.Properties, isStorageSecondaryKeyInUse)
			cf.ExtendedAuditingPolicy.RetentionInDays = convert.ToFloat64(res.Properties, retentionDays)
		}
	}
	return cf
}
