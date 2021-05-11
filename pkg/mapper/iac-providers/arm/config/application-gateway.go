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

const wafConfiguration = "webApplicationFirewallConfiguration"

// applicationGatewayConfig holds config for azurerm_application_gateway
type applicationGatewayConfig struct {
	config
	Configuration struct {
		WafConfiguration struct {
			Enabled bool `json:"enabled"`
		}
	} `json:"waf_configuration"`
}

// ApplicationGatewayConfig returns config for azurerm_application_gateway
func ApplicationGatewayConfig(r types.Resource, params map[string]interface{}) interface{} {
	cf := applicationGatewayConfig{
		config: config{
			Name:     fn.LookUp(nil, params, r.Name).(string),
			Location: fn.LookUp(nil, params, location).(string),
		},
	}
	wafc := convert.ToMap(r.Properties, wafConfiguration)
	cf.Configuration.WafConfiguration.Enabled = convert.ToBool(wafc, enabled)
	return cf
}
