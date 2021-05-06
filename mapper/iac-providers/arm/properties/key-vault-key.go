package properties

// KeyVaultKey exposes the properties for key-vault-key resource.
var KeyVaultKey *keyVaultKey

type keyVaultKey struct {
	Attributes string
}

func init() {
	KeyVaultKey = &keyVaultKey{
		Attributes: "attributes",
	}
}
