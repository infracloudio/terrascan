package config

import (
	"github.com/accurics/terrascan/mapper/convert"
	fn "github.com/accurics/terrascan/mapper/iac-providers/arm/functions"
	prop "github.com/accurics/terrascan/mapper/iac-providers/arm/properties"
	"github.com/accurics/terrascan/mapper/iac-providers/arm/template"
)

// containerRegistryConfig holds config for azurerm_container_registry
type containerRegistryConfig struct {
	config

	SKU          string `json:"sku"`
	AdminEnabled bool   `json:"admin_enabled"`
}

// ContainerRegistryConfig returns config for azurerm_container_registry
func ContainerRegistryConfig(r template.Resource, params map[string]interface{}) interface{} {
	cf := containerRegistryConfig{
		config: config{
			Location: convert.ToString(params, prop.Resource.Location),
			Name:     fn.LookUp(nil, params, r.Name),
		},
		SKU: fn.LookUp(nil, params, r.SKU.Name),
	}

	key, err := fn.Parameters(convert.ToString(r.Properties, "adminUserEnabled"))
	if err == nil {
		cf.AdminEnabled = convert.ToBool(params, key)
	}
	return cf
}
