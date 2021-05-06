package config

import (
	"github.com/accurics/terrascan/mapper/convert"
	fn "github.com/accurics/terrascan/mapper/iac-providers/arm/functions"
	prop "github.com/accurics/terrascan/mapper/iac-providers/arm/properties"
	"github.com/accurics/terrascan/mapper/iac-providers/arm/template"
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
func RoleAssignmentConfig(r template.Resource, vars, params map[string]interface{}) interface{} {
	return roleAssignmentConfig{
		config: config{
			Location: convert.ToString(params, prop.Resource.Location),
			Name:     fn.LookUp(nil, params, r.Name),
		},
		Scope:            fn.LookUp(vars, params, r.Scope),
		LockLevel:        convert.ToString(r.Properties, prop.RoleAssignment.Level),
		PrincipalID:      convert.ToString(r.Properties, prop.RoleAssignment.PrincipalID),
		RoleDefinitionID: fn.LookUp(vars, params, convert.ToString(r.Properties, prop.RoleAssignment.RoleDefinitionID)),
	}
}
