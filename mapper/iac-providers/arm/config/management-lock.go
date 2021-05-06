package config

import (
	"github.com/accurics/terrascan/mapper/convert"
	fn "github.com/accurics/terrascan/mapper/iac-providers/arm/functions"
	prop "github.com/accurics/terrascan/mapper/iac-providers/arm/properties"
	"github.com/accurics/terrascan/mapper/iac-providers/arm/template"
)

// managementLockConfig holds config for azurerm_management_lock
type managementLockConfig struct {
	config

	Scope     string `json:"scope"`
	LockLevel string `json:"lock_level"`
	Notes     string `json:"notes"`
}

// ManagementLockConfig returns config for azurerm_management_lock
func ManagementLockConfig(r template.Resource, vars, params map[string]interface{}) interface{} {
	return managementLockConfig{
		config: config{
			Location: convert.ToString(params, prop.Resource.Location),
			Name:     fn.LookUp(nil, params, r.Name),
		},
		Scope:     fn.LookUp(vars, params, r.Scope),
		LockLevel: convert.ToString(r.Properties, prop.ManagementLock.Level),
		Notes:     convert.ToString(r.Properties, prop.ManagementLock.Notes),
	}
}
