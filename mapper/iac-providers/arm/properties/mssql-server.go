package properties

// MSSQLServer exposes the properties for azurerm_mssql_server resource.
var MSSQLServer *mssqlServer

type mssqlServer struct {
	AdministratorLogin         string
	AdministratorLoginPassword string
	MinimumTLSVersion          string

	AuditingSettings           string
	StorageEndpoint            string
	StorageAccountAccessKey    string
	IsStorageSecondaryKeyInUse string
	RetentionDays              string
}

func init() {
	MSSQLServer = &mssqlServer{
		AdministratorLogin:         "administratorLogin",
		AdministratorLoginPassword: "administratorLoginPassword",
		MinimumTLSVersion:          "minimalTlsVersion",
		AuditingSettings:           "auditingSettings",
		StorageEndpoint:            "storageEndpoint",
		StorageAccountAccessKey:    "storageAccountAccessKey",
		IsStorageSecondaryKeyInUse: "isStorageSecondaryKeyInUse",
		RetentionDays:              "retentionDays",
	}
}
