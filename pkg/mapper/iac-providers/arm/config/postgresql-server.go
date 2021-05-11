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
	statusEnabled = "ENABLED"
)

// postgreSQLServerConfig holds config for azurerm_postgresql_server
type postgreSQLServerConfig struct {
	config

	SKUName                   string  `json:"sku_name"`
	StorageMB                 float64 `json:"storage_mb"`
	Version                   string  `json:"version"`
	BackupRetentionDays       float64 `json:"backup_retention_days"`
	GeoRedundantBackupEnabled bool    `json:"geo_redundant_backup_enabled"`
	SSLEnforcementEnabled     bool    `json:"ssl_enforcement_enabled"`
}

// PostgreSQLServerConfig returns config for azurerm_postgresql_server
func PostgreSQLServerConfig(r types.Resource, vars, params map[string]interface{}) interface{} {
	cf := postgreSQLServerConfig{
		config: config{
			Location: fn.LookUp(nil, params, r.Location).(string),
			Name:     fn.LookUp(nil, params, r.Name).(string),
			Tags:     r.Tags,
		},
		SKUName: fn.LookUp(vars, params, r.SKU.Name).(string),
		Version: fn.LookUp(vars, params, convert.ToString(r.Properties, version)).(string),
	}

	if profile := convert.ToMap(r.Properties, storageProfile); profile != nil {
		status := fn.LookUp(vars, params, convert.ToString(profile, geoRedundantBackup))
		cf.GeoRedundantBackupEnabled = strings.EqualFold(strings.ToUpper(status.(string)), statusEnabled)

		value := fn.LookUp(vars, params, convert.ToString(profile, backupRetentionDays))
		cf.BackupRetentionDays = value.(float64)

		value = fn.LookUp(vars, params, convert.ToString(profile, storageMB))
		cf.StorageMB = value.(float64)

		status = fn.LookUp(vars, params, convert.ToString(profile, sslEnforcement))
		cf.SSLEnforcementEnabled = strings.EqualFold(strings.ToUpper(status.(string)), statusEnabled)
	}
	return cf
}
