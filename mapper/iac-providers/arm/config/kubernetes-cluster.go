package config

import (
	"github.com/accurics/terrascan/mapper/convert"
	fn "github.com/accurics/terrascan/mapper/iac-providers/arm/functions"
	prop "github.com/accurics/terrascan/mapper/iac-providers/arm/properties"
	"github.com/accurics/terrascan/mapper/iac-providers/arm/template"
)

// kubernetesClusterConfig holds config for azurerm_kubernetes_cluster
type kubernetesClusterConfig struct {
	config

	DNSPrefix       string                  `json:"dns_prefix"`
	DefaultNodePool []pool                  `json:"default_node_pool"`
	AddonProfiles   map[string]addonProfile `json:"addon_profile"`
	NetworkProfile  struct {
		NetworkPlugin string `json:"network_plugin"`
		NetworkPolicy string `json:"network_policy"`
	} `json:"network_profile"`
}

type pool struct {
	Name      string  `json:"name"`
	NodeCount float64 `json:"node_count"`
	VMSize    string  `json:"vm_size"`
}

type addonProfile struct {
	Enabled bool                   `json:"enabled"`
	Config  map[string]interface{} `json:"config"`
}

// KubernetesClusterConfig returns config for azurerm_kubernetes_cluster.
func KubernetesClusterConfig(r template.Resource, vars, params map[string]interface{}) interface{} {
	cf := kubernetesClusterConfig{
		config: config{
			Location: convert.ToString(params, prop.Resource.Location),
			Name:     fn.LookUp(vars, params, r.Name),
		},
		DNSPrefix:     fn.LookUp(vars, params, convert.ToString(r.Properties, prop.KubernetesCluster.DNSPrefix)),
		AddonProfiles: make(map[string]addonProfile),
	}

	poolProfiles := convert.ToSlice(r.Properties, prop.KubernetesCluster.AgentPoolProfiles)
	for _, p := range poolProfiles {
		profile := p.(map[string]interface{})
		key := convert.ToString(profile, prop.KubernetesCluster.NodeCount)
		key, _ = fn.Parameters(key)

		newPool := pool{
			Name:      fn.LookUp(vars, params, convert.ToString(profile, prop.KubernetesCluster.PoolName)),
			VMSize:    fn.LookUp(vars, params, convert.ToString(profile, prop.KubernetesCluster.VMSize)),
			NodeCount: convert.ToFloat64(params, key),
		}
		cf.DefaultNodePool = append(cf.DefaultNodePool, newPool)
	}

	addonProfiles := convert.ToMap(r.Properties, prop.KubernetesCluster.AddonProfiles)
	for key, value := range addonProfiles {
		addon := value.(map[string]interface{})
		profile := addonProfile{
			Enabled: addon["enabled"].(bool),
		}

		if cfg, ok := addon["config"]; ok {
			profile.Config = cfg.(map[string]interface{})
		}

		if key == "kubeDashboard" {
			cf.AddonProfiles["kube_dashboard"] = profile
		}
		cf.AddonProfiles[key] = profile
	}

	netProfile := convert.ToMap(r.Properties, prop.KubernetesCluster.NetworkProfile)
	cf.NetworkProfile.NetworkPlugin = fn.LookUp(vars, params, convert.ToString(netProfile, prop.KubernetesCluster.NetworkPlugin))
	cf.NetworkProfile.NetworkPolicy = fn.LookUp(vars, params, convert.ToString(netProfile, prop.KubernetesCluster.NetworkPolicy))

	return cf
}
