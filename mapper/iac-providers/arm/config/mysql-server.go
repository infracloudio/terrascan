package config

import (
	"strings"

	"github.com/accurics/terrascan/mapper/convert"
	fn "github.com/accurics/terrascan/mapper/iac-providers/arm/functions"
	prop "github.com/accurics/terrascan/mapper/iac-providers/arm/properties"
	"github.com/accurics/terrascan/mapper/iac-providers/arm/template"
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
func MySQLServerConfig(r template.Resource, vars, params map[string]interface{}) interface{} {
	const enabled = "ENABLED"
	cf := mySQLServerConfig{
		config: config{
			Location: convert.ToString(params, prop.Resource.Location),
			Name:     fn.LookUp(vars, params, r.Name),
		},
		SKUName:                    fn.LookUp(vars, params, r.SKU.Name),
		Version:                    fn.LookUp(vars, params, convert.ToString(r.Properties, prop.MySQLServer.Version)),
		AdministratorLogin:         fn.LookUp(vars, params, convert.ToString(r.Properties, prop.MySQLServer.AdministratorLogin)),
		AdministratorLoginPassword: fn.LookUp(vars, params, convert.ToString(r.Properties, prop.MySQLServer.AdministratorLoginPassword)),
	}

	profile := convert.ToMap(r.Properties, prop.MySQLServer.StorageProfile)
	key, _ := fn.Parameters(convert.ToString(profile, prop.MySQLServer.StorageMB))
	cf.StorageMB = convert.ToFloat64(params, key)

	cf.BackupRetentionDays = convert.ToFloat64(profile, prop.MySQLServer.BackupRetentionDays)

	status := strings.ToUpper(convert.ToString(profile, prop.MySQLServer.GeoRedundantBackup))
	cf.GeoRedundantBackupEnabled = strings.EqualFold(status, enabled)

	status = strings.ToUpper(convert.ToString(r.Properties, prop.MySQLServer.SSLEnforcement))
	cf.SSLEnforcementEnabled = strings.EqualFold(status, enabled)
	return cf
}
