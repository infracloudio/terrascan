package properties

// MySQLServer exposes the properties for mysql-server resource.
var MySQLServer *mySQLServer

type mySQLServer struct {
	Version                    string
	AdministratorLogin         string
	AdministratorLoginPassword string
	StorageProfile             string
	StorageMB                  string
	BackupRetentionDays        string
	GeoRedundantBackup         string
	SSLEnforcement             string
}

func init() {
	MySQLServer = &mySQLServer{
		Version:                    "version",
		AdministratorLogin:         "administratorLogin",
		AdministratorLoginPassword: "administratorLoginPassword",
		StorageProfile:             "storageProfile",
		StorageMB:                  "storageMB",
		BackupRetentionDays:        "backupRetentionDays",
		GeoRedundantBackup:         "geoRedundantBackup",
		SSLEnforcement:             "sslEnforcement",
	}
}
