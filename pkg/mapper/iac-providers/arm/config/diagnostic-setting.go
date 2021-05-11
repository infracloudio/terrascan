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

const (
	storageAccountID = "storageAccountId"
	category         = "category"
	logs             = "logs"
)

// diagnosticSettingConfig holds config for azurerm_monitor_diagnostic_setting
type diagnosticSettingConfig struct {
	config

	TargetResourceID string `json:"target_resource_id"`
	StorageAccountID string `json:"storage_account_id"`
	Log              []log  `json:"log"`
}

type log struct {
	Category        string `json:"category"`
	Enabled         bool   `json:"enabled"`
	RetentionPolicy struct {
		Enabled bool    `json:"enabled"`
		Days    float64 `json:"days"`
	} `json:"retention_policy"`
}

// DiagnosticSettingConfig returns config for azurerm_monitor_diagnostic_setting
func DiagnosticSettingConfig(r types.Resource, vars, params map[string]interface{}) interface{} {
	cf := diagnosticSettingConfig{
		config: config{
			Location: fn.LookUp(nil, params, r.Location).(string),
			Name:     fn.LookUp(nil, params, r.Name).(string),
		},
		TargetResourceID: fn.LookUp(vars, params, getTargetResourceID(r.DependsOn)).(string),
		StorageAccountID: fn.LookUp(vars, params, convert.ToString(r.Properties, storageAccountID)).(string),
	}

	logs := convert.ToSlice(r.Properties, logs)
	if len(logs) > 0 {
		for _, lg := range logs {
			mp := lg.(map[string]interface{})
			policy := convert.ToMap(mp, retentionPolicy)

			l := log{
				Enabled:  convert.ToBool(mp, enabled),
				Category: convert.ToString(mp, category),
			}

			l.RetentionPolicy.Enabled = convert.ToBool(policy, enabled)
			if l.RetentionPolicy.Enabled {
				l.RetentionPolicy.Days = fn.LookUp(vars, params, convert.ToString(policy, days)).(float64)
			}
			cf.Log = append(cf.Log, l)
		}
	}
	return cf
}

// TODO: This needs to be double checked
func getTargetResourceID(deps []string) string {
	for _, d := range deps {
		if strings.Contains(d, "vault") {
			return d
		}
	}
	return ""
}
