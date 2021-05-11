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
	startIP = "startIP"
	endIP   = "endIP"
)

// redisFirewallRuleConfig holds config for azurerm_redis_firewall_rule
type redisFirewallRuleConfig struct {
	config

	StartIP string `json:"start_ip"`
	EndIP   string `json:"end_ip"`
}

// RedisFirewallRuleConfig returns config for azurerm_redis_firewall_rule
func RedisFirewallRuleConfig(r types.Resource, params map[string]interface{}) interface{} {
	return redisFirewallRuleConfig{
		config: config{
			Location: fn.LookUp(nil, params, r.Location).(string),
			Name:     fn.LookUp(nil, params, r.Name).(string),
			Tags:     r.Tags,
		},
		StartIP: fn.LookUp(nil, params, convert.ToString(r.Properties, startIP)).(string),
		EndIP:   fn.LookUp(nil, params, convert.ToString(r.Properties, endIP)).(string),
	}
}
