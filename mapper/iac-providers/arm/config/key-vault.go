package config

import (
	"github.com/accurics/terrascan/mapper/convert"
	fn "github.com/accurics/terrascan/mapper/iac-providers/arm/functions"
	prop "github.com/accurics/terrascan/mapper/iac-providers/arm/properties"
	"github.com/accurics/terrascan/mapper/iac-providers/arm/template"
)

// keyVaultConfig holds config for azurerm_key_vault
type keyVaultConfig struct {
	config

	ResourceGroupName        string `json:"resource_group_name"`
	EnabledForDiskEncryption bool   `json:"enabled_for_disk_encryption"`
	TenantID                 string `json:"tenant_id"`
	SoftDeleteEnabled        bool   `json:"soft_delete_enabled"`
}

// KeyVaultConfig returns config for azurerm_key_vault
func KeyVaultConfig(r template.Resource, params map[string]interface{}) interface{} {
	return keyVaultConfig{
		config: config{
			Location: convert.ToString(params, prop.Resource.Location),
			Name:     fn.LookUp(nil, params, r.Name),
			Tags:     nil,
		},
		SoftDeleteEnabled: convert.ToBool(r.Properties, prop.KeyVault.EnableSoftDelete),
		TenantID:          convert.ToString(params, prop.KeyVault.TenantID),
	}
}
