package properties

// ManagedDisk exposes the properties for managed-disk resource.
var ManagedDisk *managedDisk

type managedDisk struct {
	EncryptionSettingsCollection string

	CreationData     string
	CreateOption     string
	DiskSizeGB       string
	SourceResourceID string
}

func init() {
	ManagedDisk = &managedDisk{
		EncryptionSettingsCollection: "encryptionSettingsCollection",
		CreationData:                 "creationData",
		CreateOption:                 "createOption",
		DiskSizeGB:                   "diskSizeGB",
		SourceResourceID:             "sourceResourceId",
	}
}
