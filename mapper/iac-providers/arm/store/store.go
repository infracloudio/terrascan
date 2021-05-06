package store

// ResourceTypes holds mapping for ARM resource types to TF types
var ResourceTypes = map[string]string{
	"Microsoft.KeyVault/vaults":                        AzureRMKeyVault,
	"Microsoft.KeyVault/vaults/keys":                   AzureRMKeyVaultKey,
	"Microsoft.KeyVault/vaults/secrets":                AzureRMKeyVaultSecret,
	"Microsoft.Network/applicationGateways":            AzureRMApplicationGateway,
	"Microsoft.ApiManagement/service":                  AzureRMAPIManagement,
	"Microsoft.ManagedIdentity/userAssignedIdentities": AzureRMUserAssignedIdentity,
	"Microsoft.Insights/diagnosticsettings":            AzureRMMonitorDiagnosticSetting,
	"Microsoft.ContainerService/managedClusters":       AzureRMKubernetesCluster,
	"Microsoft.Compute/disks":                          AzureRMManagedDisk,
	"Microsoft.DocumentDB/databaseAccounts":            AzureRMCosmosDBAccount,
	"Microsoft.ContainerRegistry/registries":           AzureRMContainerRegistry,
	"Microsoft.Authorization/locks":                    AzureRMManagementLock,
	"Microsoft.Authorization/roleAssignments":          AzureRMRoleAssignment,
	"Microsoft.Sql/servers":                            AzureRMMSSQLServer,
	"Microsoft.DBforMySQL/servers":                     AzureRMMySQLServer,
	"Microsoft.Network/networkWatchers/flowLogs":       AzureRMNetworkWatcherFlowLog,
}
