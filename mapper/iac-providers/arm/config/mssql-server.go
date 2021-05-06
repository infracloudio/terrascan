package config

import (
	"strings"

	"github.com/accurics/terrascan/mapper/convert"
	fn "github.com/accurics/terrascan/mapper/iac-providers/arm/functions"
	prop "github.com/accurics/terrascan/mapper/iac-providers/arm/properties"
	"github.com/accurics/terrascan/mapper/iac-providers/arm/template"
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
func MSSQLServerConfig(r template.Resource, vars, params map[string]interface{}) interface{} {
	cf := mssqlServerConfig{
		config: config{
			Location: convert.ToString(params, prop.Resource.Location),
			Name:     fn.LookUp(nil, params, r.Name),
		},
	}

	key, _ := fn.Parameters(convert.ToString(r.Properties, prop.MSSQLServer.AdministratorLogin))
	cf.AdministratorLogin = convert.ToString(params, key)

	key, _ = fn.Parameters(convert.ToString(r.Properties, prop.MSSQLServer.AdministratorLoginPassword))
	cf.AdministratorLoginPassword = convert.ToString(params, key)

	key, _ = fn.Parameters(convert.ToString(r.Properties, prop.MSSQLServer.MinimumTLSVersion))
	cf.MinimumTLSVersion = convert.ToString(params, key)

	for _, res := range r.Resources {
		if strings.EqualFold(res.Type, prop.MSSQLServer.AuditingSettings) {
			cf.ExtendedAuditingPolicy.StorageEndpoint = convert.ToString(res.Properties, prop.MSSQLServer.StorageEndpoint)
			cf.ExtendedAuditingPolicy.StorageAccountAccessKey = convert.ToString(res.Properties, prop.MSSQLServer.StorageAccountAccessKey)
			cf.ExtendedAuditingPolicy.StorageAccountAccessKeyIsSecondary = convert.ToBool(res.Properties, prop.MSSQLServer.IsStorageSecondaryKeyInUse)
			cf.ExtendedAuditingPolicy.RetentionInDays = convert.ToFloat64(res.Properties, prop.MSSQLServer.RetentionDays)
		}
	}
	return cf
}
