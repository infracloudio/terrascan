package config

import (
	"time"

	"github.com/accurics/terrascan/mapper/convert"
	fn "github.com/accurics/terrascan/mapper/iac-providers/arm/functions"
	prop "github.com/accurics/terrascan/mapper/iac-providers/arm/properties"
	"github.com/accurics/terrascan/mapper/iac-providers/arm/template"
)

// keyVaultSecretConfig holds config for azurerm_key_vault_secret
type keyVaultSecretConfig struct {
	config
	ExpirationDate string `json:"expiration_date"`
}

// KeyVaultSecretConfig returns config for azurerm_key_vault_secret
func KeyVaultSecretConfig(r template.Resource, params map[string]interface{}) interface{} {
	a := convert.ToMap(r.Properties, prop.KeyVaultSecret.Attributes)
	cf := keyVaultSecretConfig{
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
