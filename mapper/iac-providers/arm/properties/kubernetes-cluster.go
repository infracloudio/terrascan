package properties

// KubernetesCluster exposes the properties for azurerm_kubernetes_cluster resource.
var KubernetesCluster *kubernetesCluster

type kubernetesCluster struct {
	DNSPrefix     string
	AddonProfiles string

	AgentPoolProfiles string
	PoolName          string
	NodeCount         string
	VMSize            string

	NetworkProfile string
	NetworkPlugin  string
	NetworkPolicy  string
}

func init() {
	KubernetesCluster = &kubernetesCluster{
		DNSPrefix:         "dnsPrefix",
		AgentPoolProfiles: "agentPoolProfiles",
		PoolName:          "name",
		NodeCount:         "count",
		VMSize:            "vmSize",
		AddonProfiles:     "addonProfiles",
		NetworkProfile:    "networkProfile",
		NetworkPlugin:     "networkPlugin",
		NetworkPolicy:     "networkPolicy",
	}
}
