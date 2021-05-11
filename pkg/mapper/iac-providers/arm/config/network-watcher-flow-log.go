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
	targetResourceID                         = "targetResourceId"
	storageID                                = "storageId"
	enabled                                  = "enabled"
	retentionPolicy                          = "retentionPolicy"
	days                                     = "days"
	flowAnalyticsConfiguration               = "flowAnalyticsConfiguration"
	networkWatcherFlowAnalyticsConfiguration = "networkWatcherFlowAnalyticsConfiguration"
	workspaceID                              = "workspaceId"
	workspaceRegion                          = "workspaceRegion"
	workspaceResourceID                      = "workspaceResourceId"
	trafficAnalyticsInterval                 = "trafficAnalyticsInterval"
)

// networkWatcherFlowLogConfig holds config for azurerm_network_watcher_flow_log
type networkWatcherFlowLogConfig struct {
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

// NetworkWatcherFlowLogConfig returns config for azurerm_network_watcher_flow_log
func NetworkWatcherFlowLogConfig(r types.Resource, vars, params map[string]interface{}) interface{} {
	cf := networkWatcherFlowLogConfig{
		config: config{
			Location: fn.LookUp(vars, params, r.Location).(string),
			Name:     fn.LookUp(vars, params, r.Name).(string),
		},
		NetworkSecurityGroupID: fn.LookUp(vars, params, convert.ToString(r.Properties, targetResourceID)).(string),
		StorageAccountID:       fn.LookUp(vars, params, convert.ToString(r.Properties, storageID)).(string),
		Enabled:                convert.ToBool(r.Properties, enabled),
	}

	policy := convert.ToMap(r.Properties, retentionPolicy)
	cf.RetentionPolicy.Enabled = convert.ToBool(policy, enabled)
	cf.RetentionPolicy.Days = fn.LookUp(vars, params, convert.ToString(policy, days)).(float64)

	flowConfig := convert.ToMap(r.Properties, flowAnalyticsConfiguration)
	if flowConfig != nil {
		networkConfig := convert.ToMap(flowConfig, networkWatcherFlowAnalyticsConfiguration)
		cf.TrafficAnalytics.Enabled = convert.ToBool(networkConfig, enabled)
		cf.TrafficAnalytics.WorkspaceID = fn.LookUp(vars, params, workspaceID).(string)
		cf.TrafficAnalytics.WorkspaceRegion = fn.LookUp(vars, params, workspaceRegion).(string)
		cf.TrafficAnalytics.WorkspaceResourceID = fn.LookUp(vars, params, workspaceResourceID).(string)
		cf.TrafficAnalytics.IntervalInMinutes = convert.ToFloat64(networkConfig, trafficAnalyticsInterval)
	}

	return cf
}
