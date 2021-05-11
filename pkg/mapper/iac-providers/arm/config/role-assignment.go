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
	level            = "level"
	principalID      = "principalId"
	roleDefinitionID = "roleDefinitionId"
)

// roleAssignmentConfig holds config for azurerm_role_assignment
type roleAssignmentConfig struct {
	config

	Scope            string `json:"scope"`
	LockLevel        string `json:"lock_level"`
	PrincipalID      string `json:"principal_id"`
	RoleDefinitionID string `json:"role_definition_id"`
}

// RoleAssignmentConfig returns config for azurerm_role_assignment
func RoleAssignmentConfig(r types.Resource, vars, params map[string]interface{}) interface{} {
	return roleAssignmentConfig{
		config: config{
			Location: fn.LookUp(nil, params, r.Location).(string),
			Name:     fn.LookUp(nil, params, r.Name).(string),
		},
		Scope:            fn.LookUp(vars, params, r.Scope).(string),
		LockLevel:        convert.ToString(r.Properties, level),
		PrincipalID:      convert.ToString(r.Properties, principalID),
		RoleDefinitionID: fn.LookUp(vars, params, convert.ToString(r.Properties, roleDefinitionID)).(string),
	}
}
