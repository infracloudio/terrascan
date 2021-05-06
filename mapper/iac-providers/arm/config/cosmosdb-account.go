package config

import (
	"fmt"

	"github.com/accurics/terrascan/mapper/convert"
	fn "github.com/accurics/terrascan/mapper/iac-providers/arm/functions"
	prop "github.com/accurics/terrascan/mapper/iac-providers/arm/properties"
	"github.com/accurics/terrascan/mapper/iac-providers/arm/template"
)

// cosmosDBAccountConfig holds config for azurerm_cosmosdb_account
type cosmosDBAccountConfig struct {
	config

	IPRangeFilter interface{} `json:"ip_range_filter"`
	Tags          interface{} `json:"tags"`
}

// CosmosDBAccountConfig returns config for azurerm_cosmosdb_account
func CosmosDBAccountConfig(r template.Resource, params map[string]interface{}) interface{} {
	var v interface{}
	ipr := convert.ToSlice(r.Properties, prop.CosmosDBAccount.IPRules)
	for _, s := range ipr {
		m := s.(map[string]interface{})
		v = convert.ToString(m, prop.CosmosDBAccount.IPAddressOrRange)
		break
	}
	cf := cosmosDBAccountConfig{
		config: config{
			Location: convert.ToString(params, prop.Resource.Location),
			Name:     fn.LookUp(nil, params, r.Name),
		},
	}
	if v != nil && len(v.(string)) > 0 {
		cf.IPRangeFilter = v
	}
	if r.Tags != nil && len(r.Tags) > 0 {
		cf.Tags = fmt.Sprintf("%v", r.Tags)
	}
	return cf
}
