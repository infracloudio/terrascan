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
	emails              = "emails"
	phone               = "phone"
	alertNotifications  = "alertNotifications"
	notificationsByRole = "notificationsByRole"
	state               = "state"
)

// securityCenterContactConfig holds config for azurerm_security_center_contact
type securityCenterContactConfig struct {
	config

	Email string `json:"email"`
	Phone string `json:"phone"`

	AlertNotifications bool `json:"alert_notifications"`
	AlertsToAdmins     bool `json:"alerts_to_admins"`
}

// SecurityCenterContactConfig returns config for azurerm_security_center_contact
func SecurityCenterContactConfig(r types.Resource, params map[string]interface{}) interface{} {
	cf := securityCenterContactConfig{
		config: config{
			Location: fn.LookUp(nil, params, r.Location).(string),
			Name:     fn.LookUp(nil, params, r.Name).(string),
		},
		Phone: fn.LookUp(nil, params, convert.ToString(r.Properties, phone)).(string),
		Email: fn.LookUp(nil, params, convert.ToString(r.Properties, emails)).(string),
	}

	notifications := convert.ToMap(r.Properties, alertNotifications)
	state := convert.ToString(notifications, state)
	cf.AlertNotifications = strings.EqualFold(strings.ToUpper(state), "ON")

	notifications = convert.ToMap(r.Properties, notificationsByRole)
	state = convert.ToString(notifications, state)
	cf.AlertNotifications = strings.EqualFold(strings.ToUpper(state), "ON")

	return cf
}
