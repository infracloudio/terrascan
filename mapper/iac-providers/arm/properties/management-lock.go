package properties

// ManagementLock exposes the properties for azurerm_management_lock resource.
var ManagementLock *managementLock

type managementLock struct {
	Level string
	Notes string
}

func init() {
	ManagementLock = &managementLock{
		Level: "level",
		Notes: "notes",
	}
}
