package config

import (
	"github.com/accurics/terrascan/mapper/convert"
	fn "github.com/accurics/terrascan/mapper/iac-providers/arm/functions"
	prop "github.com/accurics/terrascan/mapper/iac-providers/arm/properties"
	"github.com/accurics/terrascan/mapper/iac-providers/arm/template"
)

// managedDiskConfig holds config for azurerm_managed_disk
type managedDiskConfig struct {
	config

	CreateOption       string                 `json:"create_option"`
	DiskSizeGB         float64                `json:"disk_size_gb"`
	SourceResourceID   string                 `json:"source_resource_id"`
	StorageAccountType string                 `json:"storage_account_type"`
	EncryptionSettings map[string]interface{} `json:"encryption_settings"`
}

// ManagedDiskConfig returns config for azurerm_managed_disk.
func ManagedDiskConfig(r template.Resource, vars, params map[string]interface{}) interface{} {
	cf := managedDiskConfig{
		config: config{
			Location: convert.ToString(params, prop.Resource.Location),
			Name:     fn.LookUp(vars, params, r.Name),
		},
		StorageAccountType: r.SKU.Name,
		EncryptionSettings: convert.ToMap(r.Properties, prop.ManagedDisk.EncryptionSettingsCollection),
	}

	data := convert.ToMap(r.Properties, prop.ManagedDisk.CreationData)
	cf.CreateOption = convert.ToString(data, prop.ManagedDisk.CreateOption)
	cf.SourceResourceID = convert.ToString(data, prop.ManagedDisk.SourceResourceID)

	key, _ := fn.Parameters(convert.ToString(r.Properties, prop.ManagedDisk.DiskSizeGB))
	cf.DiskSizeGB = convert.ToFloat64(params, key)
	return cf
}
