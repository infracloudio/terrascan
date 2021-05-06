package config

import (
	"github.com/accurics/terrascan/mapper/convert"
	fn "github.com/accurics/terrascan/mapper/iac-providers/arm/functions"
	prop "github.com/accurics/terrascan/mapper/iac-providers/arm/properties"
	"github.com/accurics/terrascan/mapper/iac-providers/arm/template"
)

// applicationGatewayConfig holds config for azurerm_application_gateway
type applicationGatewayConfig struct {
	config
	Configuration struct {
		WafConfiguration struct {
			Enabled bool `json:"enabled"`
		}
	} `json:"waf_configuration"`
}

// ApplicationGatewayConfig returns config for azurerm_application_gateway
func ApplicationGatewayConfig(r template.Resource, params map[string]interface{}) interface{} {
	cf := applicationGatewayConfig{
		config: config{
			Name:     fn.LookUp(nil, params, r.Name),
			Location: convert.ToString(params, prop.Resource.Location),
		},
	}
	wafc := convert.ToMap(r.Properties, prop.ApplicationGateway.WAFConfiguration)
	cf.Configuration.WafConfiguration.Enabled = convert.ToBool(wafc, prop.ApplicationGateway.Enabled)
	return cf
}
