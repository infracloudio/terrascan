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
	version             = "version"
	storageProfile      = "storageProfile"
	storageMB           = "storageMB"
	backupRetentionDays = "backupRetentionDays"
	geoRedundantBackup  = "geoRedundantBackup"
	sslEnforcement      = "sslEnforcement"
)

// mySQLServerConfig holds config for azurerm_mysql_server
type mySQLServerConfig struct {
	config

	AdministratorLogin         string  `json:"administrator_login"`
	AdministratorLoginPassword string  `json:"administrator_login_password"`
	SKUName                    string  `json:"sku_name"`
	StorageMB                  float64 `json:"storage_mb"`
	Version                    string  `json:"version"`
	BackupRetentionDays        float64 `json:"backup_retention_days"`
	GeoRedundantBackupEnabled  bool    `json:"geo_redundant_backup_enabled"`
	SSLEnforcementEnabled      bool    `json:"ssl_enforcement_enabled"`
}

// MySQLServerConfig returns config for azurerm_mysql_server
func MySQLServerConfig(r types.Resource, vars, params map[string]interface{}) interface{} {
	const enabled = "ENABLED"
	cf := mySQLServerConfig{
		config: config{
			Location: fn.LookUp(nil, params, r.Location).(string),
			Name:     fn.LookUp(nil, params, r.Name).(string),
		},
		SKUName:                    fn.LookUp(vars, params, r.SKU.Name).(string),
		Version:                    fn.LookUp(vars, params, convert.ToString(r.Properties, version)).(string),
		AdministratorLogin:         fn.LookUp(vars, params, convert.ToString(r.Properties, administratorLogin)).(string),
		AdministratorLoginPassword: fn.LookUp(vars, params, convert.ToString(r.Properties, administratorLoginPassword)).(string),
	}

	profile := convert.ToMap(r.Properties, storageProfile)
	cf.StorageMB = fn.LookUp(vars, params, convert.ToString(profile, storageMB)).(float64)

	cf.BackupRetentionDays = convert.ToFloat64(profile, backupRetentionDays)

	status := strings.ToUpper(convert.ToString(profile, geoRedundantBackup))
	cf.GeoRedundantBackupEnabled = strings.EqualFold(status, enabled)

	status = strings.ToUpper(convert.ToString(r.Properties, sslEnforcement))
	cf.SSLEnforcementEnabled = strings.EqualFold(status, enabled)
	return cf
}
