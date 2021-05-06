package config

import (
	"time"

	"github.com/accurics/terrascan/mapper/convert"
	fn "github.com/accurics/terrascan/mapper/iac-providers/arm/functions"
	prop "github.com/accurics/terrascan/mapper/iac-providers/arm/properties"
	"github.com/accurics/terrascan/mapper/iac-providers/arm/template"
)

// keyVaultKeyConfig holds config for azurerm_key_vault_key
type keyVaultKeyConfig struct {
	config
	ExpirationDate string `json:"expiration_date"`
}

// KeyVaultKeyConfig returns config for azurerm_key_vault_key
func KeyVaultKeyConfig(r template.Resource, params map[string]interface{}) interface{} {
	a := convert.ToMap(r.Properties, prop.KeyVaultKey.Attributes)
	cf := keyVaultKeyConfig{
		config: config{
			Location: convert.ToString(params, prop.Resource.Location),
			Name:     fn.LookUp(nil, params, r.Name),
		},
	}
	if i := a["exp"]; i != nil {
		t := time.Unix(int64(i.(float64)), 0)
		cf.ExpirationDate = t.Format(time.RFC3339)
	}
	return cf
}
