package properties

// DiagnosticSetting exposes the properties for monitor_diagnostic_setting resource.
var DiagnosticSetting *diagnosticSetting

type diagnosticSetting struct {
	StorageAccountID string
	Logs             string
	Enabled          string
	Category         string
	RetentionPolicy  string
	Days             string
}

func init() {
	DiagnosticSetting = &diagnosticSetting{
		StorageAccountID: "storageAccountId",
		Logs:             "logs",
		Enabled:          "enabled",
		Category:         "category",
		RetentionPolicy:  "retentionPolicy",
		Days:             "days",
	}
}
