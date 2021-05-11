/*
    Copyright (C) 2020 Accurics, Inc.

	Licensed under the Apache License, Version 2.0 (the "License");
    you may not use this file except in compliance with the License.
    You may obtain a copy of the License at

		http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License.
*/

package config

import (
	"github.com/accurics/terrascan/pkg/mapper/convert"
	fn "github.com/accurics/terrascan/pkg/mapper/iac-providers/arm/functions"
	"github.com/accurics/terrascan/pkg/mapper/iac-providers/arm/types"
)

const (
	dnsPrefix         = "dnsPrefix"
	agentPoolProfiles = "agentPoolProfiles"
	poolName          = "name"
	nodeCount         = "count"
	vmSize            = "vmSize"
	addonProfiles     = "addonProfiles"
	networkProfile    = "networkProfile"
	networkPlugin     = "networkPlugin"
	networkPolicy     = "networkPolicy"
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
func KubernetesClusterConfig(r types.Resource, vars, params map[string]interface{}) interface{} {
	cf := kubernetesClusterConfig{
		config: config{
			Location: fn.LookUp(nil, params, r.Location).(string),
			Name:     fn.LookUp(nil, params, r.Name).(string),
		},
		DNSPrefix:     fn.LookUp(vars, params, convert.ToString(r.Properties, dnsPrefix)).(string),
		AddonProfiles: make(map[string]addonProfile),
	}

	poolProfiles := convert.ToSlice(r.Properties, agentPoolProfiles)
	for _, p := range poolProfiles {
		profile := p.(map[string]interface{})
		newPool := pool{
			Name:      fn.LookUp(vars, params, convert.ToString(profile, poolName)).(string),
			VMSize:    fn.LookUp(vars, params, convert.ToString(profile, vmSize)).(string),
			NodeCount: fn.LookUp(vars, params, convert.ToString(profile, nodeCount)).(float64),
		}
		cf.DefaultNodePool = append(cf.DefaultNodePool, newPool)
	}

	addonProfiles := convert.ToMap(r.Properties, addonProfiles)
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

	netProfile := convert.ToMap(r.Properties, networkProfile)
	cf.NetworkProfile.NetworkPlugin = fn.LookUp(vars, params, convert.ToString(netProfile, networkPlugin)).(string)
	cf.NetworkProfile.NetworkPolicy = fn.LookUp(vars, params, convert.ToString(netProfile, networkPolicy)).(string)

	return cf
}
