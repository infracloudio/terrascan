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
	"strings"

	"github.com/accurics/terrascan/pkg/mapper/convert"
	fn "github.com/accurics/terrascan/pkg/mapper/iac-providers/arm/functions"
	"github.com/accurics/terrascan/pkg/mapper/iac-providers/arm/types"
)

const publicAccess = "publicAccess"

// storageContainerConfig holds config for azurerm_storage_container
type storageContainerConfig struct {
	config

	ContainerAccessType string `json:"container_access_type"`
}

// StorageContainerConfig returns config for azurerm_storage_container
func StorageContainerConfig(r types.Resource, params map[string]interface{}) interface{} {
	cf := storageContainerConfig{
		config: config{
			Location: fn.LookUp(nil, params, r.Location).(string),
			Name:     fn.LookUp(nil, params, r.Name).(string),
			Tags:     r.Tags,
		},
	}

	access := fn.LookUp(nil, params, convert.ToString(r.Properties, publicAccess)).(string)
	if strings.ToUpper(access) == "NONE" {
		access = "private"
	}
	cf.ContainerAccessType = access

	return cf
}
