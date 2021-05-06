package properties

// FlowLogs exposes the properties for network_watcher_flow_log resource.
var FlowLogs *flowLogs

type flowLogs struct {
	TargetResourceID                         string
	StorageID                                string
	Enabled                                  string
	RetentionPolicy                          string
	Days                                     string
	FlowAnalyticsConfiguration               string
	NetworkWatcherFlowAnalyticsConfiguration string
	WorkspaceID                              string
	WorkspaceRegion                          string
	WorkspaceResourceID                      string
	TrafficAnalyticsInterval                 string
}

func init() {
	FlowLogs = &flowLogs{
		TargetResourceID:                         "targetResourceId",
		StorageID:                                "storageId",
		Enabled:                                  "enabled",
		RetentionPolicy:                          "retentionPolicy",
		Days:                                     "days",
		FlowAnalyticsConfiguration:               "flowAnalyticsConfiguration",
		NetworkWatcherFlowAnalyticsConfiguration: "networkWatcherFlowAnalyticsConfiguration",
		WorkspaceID:                              "workspaceId",
		WorkspaceRegion:                          "workspaceRegion",
		WorkspaceResourceID:                      "workspaceResourceId",
		TrafficAnalyticsInterval:                 "trafficAnalyticsInterval",
	}
}
