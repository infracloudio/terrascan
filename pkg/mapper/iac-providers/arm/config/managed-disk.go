/*
    Copyright (C) 2020 Accurics, Inc.

	Licensed under the Apache License, Version 2.0 (the "License");
    you may not use this file except in compliance with the License.
    You may obtain a copy of the License at

		http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License.
*/

package config

import (
	"github.com/accurics/terrascan/pkg/mapper/convert"
	fn "github.com/accurics/terrascan/pkg/mapper/iac-providers/arm/functions"
	"github.com/accurics/terrascan/pkg/mapper/iac-providers/arm/types"
)

const (
	encryptionSettingsCollection = "encryptionSettingsCollection"
	creationData                 = "creationData"
	createOption                 = "createOption"
	diskSizeGB                   = "diskSizeGB"
	sourceResourceID             = "sourceResourceId"
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
func ManagedDiskConfig(r types.Resource, vars, params map[string]interface{}) interface{} {
	cf := managedDiskConfig{
		config: config{
			Location: fn.LookUp(nil, params, r.Location).(string),
			Name:     fn.LookUp(nil, params, r.Name).(string),
		},
		StorageAccountType: r.SKU.Name,
		EncryptionSettings: convert.ToMap(r.Properties, encryptionSettingsCollection),
		DiskSizeGB:         fn.LookUp(vars, params, convert.ToString(r.Properties, diskSizeGB)).(float64),
	}

	data := convert.ToMap(r.Properties, creationData)
	cf.CreateOption = convert.ToString(data, createOption)
	cf.SourceResourceID = convert.ToString(data, sourceResourceID)

	return cf
}
