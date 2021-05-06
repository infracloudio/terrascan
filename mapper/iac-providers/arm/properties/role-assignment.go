package properties

// RoleAssignment exposes the properties for azurerm_role_assignment resource.
var RoleAssignment *roleAssignment

type roleAssignment struct {
	Level            string
	PrincipalID      string
	RoleDefinitionID string
}

func init() {
	RoleAssignment = &roleAssignment{
		Level:            "level",
		PrincipalID:      "principalId",
		RoleDefinitionID: "roleDefinitionId",
	}
}
