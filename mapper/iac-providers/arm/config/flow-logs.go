package config

import (
	"github.com/accurics/terrascan/mapper/convert"
	fn "github.com/accurics/terrascan/mapper/iac-providers/arm/functions"
	prop "github.com/accurics/terrascan/mapper/iac-providers/arm/properties"
	"github.com/accurics/terrascan/mapper/iac-providers/arm/template"
)

// flowLogsConfig holds config for azurerm_network_watcher_flow_log
type flowLogsConfig struct {
	config

	NetworkSecurityGroupID string `json:"network_security_group_id"`
	StorageAccountID       string `json:"storage_account_id"`
	Enabled                bool   `json:"enabled"`

	RetentionPolicy struct {
		Enabled bool    `json:"enabled,omitempty"`
		Days    float64 `json:"days,omitempty"`
	} `json:"retention_policy,omitempty"`

	TrafficAnalytics struct {
		Enabled             bool    `json:"enabled,omitempty"`
		WorkspaceID         string  `json:"workspace_id,omitempty"`
		WorkspaceRegion     string  `json:"workspace_region,omitempty"`
		WorkspaceResourceID string  `json:"workspace_resource_id,omitempty"`
		IntervalInMinutes   float64 `json:"interval_in_minutes,omitempty"`
	} `json:"traffic_analytics,omitempty"`
}

// FlowLogsConfig returns config for azurerm_network_watcher_flow_log
func FlowLogsConfig(r template.Resource, vars, params map[string]interface{}) interface{} {
	cf := flowLogsConfig{
		config: config{
			Location: convert.ToString(params, prop.Resource.Location),
			Name:     fn.LookUp(vars, params, r.Name),
		},
		NetworkSecurityGroupID: fn.LookUp(vars, params, convert.ToString(r.Properties, prop.FlowLogs.TargetResourceID)),
		StorageAccountID:       fn.LookUp(vars, params, convert.ToString(r.Properties, prop.FlowLogs.StorageID)),
		Enabled:                convert.ToBool(r.Properties, prop.FlowLogs.Enabled),
	}

	policy := convert.ToMap(r.Properties, prop.FlowLogs.RetentionPolicy)
	cf.RetentionPolicy.Enabled = convert.ToBool(policy, prop.FlowLogs.Enabled)
	key, _ := fn.Parameters(convert.ToString(policy, prop.FlowLogs.Days))
	cf.RetentionPolicy.Days = convert.ToFloat64(params, key)

	flowConfig := convert.ToMap(r.Properties, prop.FlowLogs.FlowAnalyticsConfiguration)
	if flowConfig != nil {
		networkConfig := convert.ToMap(flowConfig, prop.FlowLogs.NetworkWatcherFlowAnalyticsConfiguration)
		cf.TrafficAnalytics.Enabled = convert.ToBool(networkConfig, prop.FlowLogs.Enabled)
		cf.TrafficAnalytics.WorkspaceID = fn.LookUp(vars, params, prop.FlowLogs.WorkspaceID)
		cf.TrafficAnalytics.WorkspaceRegion = fn.LookUp(vars, params, prop.FlowLogs.WorkspaceRegion)
		cf.TrafficAnalytics.WorkspaceResourceID = fn.LookUp(vars, params, prop.FlowLogs.WorkspaceResourceID)
		cf.TrafficAnalytics.IntervalInMinutes = convert.ToFloat64(networkConfig, prop.FlowLogs.TrafficAnalyticsInterval)
	}

	return cf
}
