package properties

// KeyVault exposes the properties for key-vault resource.
var KeyVault *keyVault

type keyVault struct {
	EnableSoftDelete string
	TenantID         string
}

func init() {
	KeyVault = &keyVault{
		EnableSoftDelete: "enableSoftDelete",
		TenantID:         "tenantId",
	}
}
