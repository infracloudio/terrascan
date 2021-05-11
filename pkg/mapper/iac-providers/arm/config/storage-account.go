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
	supportsHTTPSTrafficOnly = "supportsHttpsTrafficOnly"
	networkAcls              = "networkAcls"
	defaultAction            = "defaultAction"
	bypass                   = "bypass"
)

// storageAccountConfig holds config for azurerm_storage_account
type storageAccountConfig struct {
	config

	AccountTier            string `json:"account_tier"`
	AccountReplicationType string `json:"account_replication_type"`
	EnableHTTPSTrafficOnly bool   `json:"enable_https_traffic_only"`

	NetworkRules struct {
		DefaultAction string   `json:"default_action"`
		ByPass        string   `json:"bypass"`
		IPRules       []string `json:"ip_rules,omitempty"`
	} `json:"network_rules,omitempty"`
}

// StorageAccountConfig returns config for azurerm_storage_account
func StorageAccountConfig(r types.Resource, vars, params map[string]interface{}) interface{} {
	cf := storageAccountConfig{
		config: config{
			Location: fn.LookUp(vars, params, r.Location).(string),
			Name:     fn.LookUp(vars, params, r.Name).(string),
			Tags:     r.Tags,
		},
		AccountTier:            r.SKU.Tier,
		AccountReplicationType: r.SKU.Name,
		EnableHTTPSTrafficOnly: convert.ToBool(r.Properties, supportsHTTPSTrafficOnly),
	}

	if acls := convert.ToMap(r.Properties, networkAcls); acls != nil {
		cf.NetworkRules.DefaultAction = fn.LookUp(vars, params, convert.ToString(acls, defaultAction)).(string)
		cf.NetworkRules.ByPass = fn.LookUp(vars, params, convert.ToString(acls, bypass)).(string)

		rules := convert.ToSlice(acls, ipRules)
		for _, rule := range rules {
			r := rule.(map[string]string)
			cf.NetworkRules.IPRules = append(cf.NetworkRules.IPRules, r[value])
		}
	}
	return cf
}
