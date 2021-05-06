package properties

// KeyVaultSecret exposes the properties for key-vault-secret resource.
var KeyVaultSecret *keyVaultSecret

type keyVaultSecret struct {
	Attributes string
}

func init() {
	KeyVaultSecret = &keyVaultSecret{
		Attributes: "attributes",
	}
}
