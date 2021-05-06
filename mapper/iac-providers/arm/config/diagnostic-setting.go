package config

import (
	"strings"

	"github.com/accurics/terrascan/mapper/convert"
	fn "github.com/accurics/terrascan/mapper/iac-providers/arm/functions"
	prop "github.com/accurics/terrascan/mapper/iac-providers/arm/properties"
	"github.com/accurics/terrascan/mapper/iac-providers/arm/template"
)

// diagnosticSettingConfig holds config for azurerm_monitor_diagnostic_setting
type diagnosticSettingConfig struct {
	config

	TargetResourceID string `json:"target_resource_id"`
	StorageAccountID string `json:"storage_account_id"`
	Log              []log  `json:"log"`
}

type log struct {
	Category        string `json:"category"`
	Enabled         bool   `json:"enabled"`
	RetentionPolicy struct {
		Enabled bool    `json:"enabled"`
		Days    float64 `json:"days"`
	} `json:"retention_policy"`
}

// DiagnosticSettingConfig returns config for azurerm_monitor_diagnostic_setting
func DiagnosticSettingConfig(r template.Resource, vars, params map[string]interface{}) interface{} {
	dsc := diagnosticSettingConfig{
		config: config{
			Location: convert.ToString(params, prop.Resource.Location),
			Name:     fn.LookUp(vars, params, r.Name),
		},
		TargetResourceID: fn.LookUp(vars, params, getTargetResourceID(r.DependsOn)),
		StorageAccountID: fn.LookUp(vars, params, convert.ToString(r.Properties, prop.DiagnosticSetting.StorageAccountID)),
	}

	logs := convert.ToSlice(r.Properties, prop.DiagnosticSetting.Logs)
	if len(logs) > 0 {
		for _, lg := range logs {
			mp := lg.(map[string]interface{})
			policy := convert.ToMap(mp, prop.DiagnosticSetting.RetentionPolicy)

			l := log{
				Enabled:  convert.ToBool(mp, prop.DiagnosticSetting.Enabled),
				Category: convert.ToString(mp, prop.DiagnosticSetting.Category),
			}

			l.RetentionPolicy.Enabled = convert.ToBool(policy, prop.DiagnosticSetting.Enabled)
			if l.RetentionPolicy.Enabled {
				key := convert.ToString(policy, prop.DiagnosticSetting.Days)
				key, _ = fn.Parameters(key)
				l.RetentionPolicy.Days = convert.ToFloat64(params, key)
			}
			dsc.Log = append(dsc.Log, l)
		}
	}
	return dsc
}

// TODO: This needs to be double checked
func getTargetResourceID(deps []string) string {
	for _, d := range deps {
		if strings.Contains(d, "vault") {
			return d
		}
	}
	return ""
}
