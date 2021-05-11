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

const pricingTier = "pricingTier"

// securityCenterSubscriptionPricingConfig holds config for azurerm_security_center_subscription_pricing
type securityCenterSubscriptionPricingConfig struct {
	config

	Tier string `json:"tier"`
}

// SecurityCenterSubscriptionPricingConfig returns config for azurerm_security_center_subscription_pricing
func SecurityCenterSubscriptionPricingConfig(r types.Resource, params map[string]interface{}) interface{} {
	return securityCenterSubscriptionPricingConfig{
		config: config{
			Location: fn.LookUp(nil, params, r.Location).(string),
			Name:     fn.LookUp(nil, params, r.Name).(string),
		},
		Tier: fn.LookUp(nil, params, convert.ToString(r.Properties, pricingTier)).(string),
	}
}
