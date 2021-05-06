package arm

import (
	"encoding/json"
	"errors"

	"github.com/accurics/terrascan/mapper/core"
	"github.com/accurics/terrascan/mapper/iac-providers/arm/config"
	fn "github.com/accurics/terrascan/mapper/iac-providers/arm/functions"
	"github.com/accurics/terrascan/mapper/iac-providers/arm/store"
	"github.com/accurics/terrascan/mapper/iac-providers/arm/template"
	"github.com/accurics/terrascan/pkg/iac-providers/output"
	"github.com/accurics/terrascan/pkg/utils"
)

const errUnsupportedDoc = "unsupported document type"

type armMapper struct {
	templateParameters map[string]interface{}
}

// Mapper returns an ARM mapper for given template schema
func Mapper() core.Mapper {
	return armMapper{}
}

// Map transforms the provider specific template to terrascan native format.
func (m armMapper) Map(doc *utils.IacDocument, params ...map[string]interface{}) (output.AllResourceConfigs, error) {
	allRC := make(map[string][]output.ResourceConfig)
	template, err := extractTemplate(doc)
	if err != nil {
		return nil, err
	}

	// set template parameters with default values if not found
	m.templateParameters = params[0]
	for key, param := range template.Parameters {
		if _, ok := m.templateParameters[key]; !ok {
			m.templateParameters[key] = param.DefaultValue
		}
	}

	// transform each resource and generate config
	for _, r := range template.Resources {
		rc := output.ResourceConfig{
			Name:      fn.LookUp(template.Variables, m.templateParameters, r.Name),
			Source:    doc.FilePath,
			Line:      doc.StartLine,
			SkipRules: nil,
		}

		// skip if resource does not have a mapping
		if t, ok := store.ResourceTypes[r.Type]; ok {
			rc.Type = t
		} else {
			continue
		}

		rc.ID = rc.Type + "." + rc.Name
		fn.ResourceIDs[r.Type] = rc.ID
		rc.Config = m.mapConfigForResource(r, template.Variables)
		allRC[rc.Type] = append(allRC[rc.Type], rc)
	}
	return allRC, nil
}

func extractTemplate(doc *utils.IacDocument) (*template.Template, error) {
	if doc.Type == utils.JSONDoc {
		var t template.Template
		err := json.Unmarshal(doc.Data, &t)
		if err != nil {
			return nil, err
		}
		return &t, nil
	}
	return nil, errors.New(errUnsupportedDoc)
}

func (m armMapper) mapConfigForResource(r template.Resource, vars map[string]interface{}) interface{} {
	switch store.ResourceTypes[r.Type] {
	case store.AzureRMKeyVault:
		return config.KeyVaultConfig(r, m.templateParameters)
	case store.AzureRMKeyVaultSecret:
		return config.KeyVaultSecretConfig(r, m.templateParameters)
	case store.AzureRMKeyVaultKey:
		return config.KeyVaultKeyConfig(r, m.templateParameters)
	case store.AzureRMApplicationGateway:
		return config.ApplicationGatewayConfig(r, m.templateParameters)
	case store.AzureRMMonitorDiagnosticSetting:
		return config.DiagnosticSettingConfig(r, vars, m.templateParameters)
	case store.AzureRMKubernetesCluster:
		return config.KubernetesClusterConfig(r, vars, m.templateParameters)
	case store.AzureRMManagedDisk:
		return config.ManagedDiskConfig(r, vars, m.templateParameters)
	case store.AzureRMCosmosDBAccount:
		return config.CosmosDBAccountConfig(r, m.templateParameters)
	case store.AzureRMContainerRegistry:
		return config.ContainerRegistryConfig(r, m.templateParameters)
	case store.AzureRMManagementLock:
		return config.ManagementLockConfig(r, vars, m.templateParameters)
	case store.AzureRMRoleAssignment:
		return config.RoleAssignmentConfig(r, vars, m.templateParameters)
	case store.AzureRMMSSQLServer:
		return config.MSSQLServerConfig(r, vars, m.templateParameters)
	case store.AzureRMMySQLServer:
		return config.MySQLServerConfig(r, vars, m.templateParameters)
	case store.AzureRMNetworkWatcherFlowLog:
		return config.FlowLogsConfig(r, vars, m.templateParameters)
	}
	return nil
}
