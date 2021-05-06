package template

// Template represents the base structure of an ARM template.
type Template struct {
	// Template schema.
	Schema string `json:"$schema"`
	// Version of the template content.
	ContentVersion string `json:"contentVersion"`
	// Template parameters.
	Parameters map[string]Parameter `json:"parameters"`
	// Template variables.
	Variables map[string]interface{} `json:"variables"`
	// A collection of Azure resources.
	Resources []Resource `json:"resources"`
	// Template output.
	Outputs interface{} `json:"outputs"`
}

// Parameter defines the values that are provided to customize
// a resource deployment.
type Parameter struct {
	// Type of parameter value.
	Type string `json:"type"`
	// Default value of parameter.
	DefaultValue interface{} `json:"defaultValue"`
	// Array of allowed values.
	AllowedValues []interface{} `json:"allowedValues"`
	// Minimum value for int.
	MinValue int `json:"minValue"`
	// Maximum value for int.
	MaxValue int `json:"maxValue"`
	// Minimum length for string or array parameter.
	MinLength int `json:"minLength"`
	// Maximum length for string or array parameter.
	MaxLength int `json:"maxLength"`
	// Metadata for parameter.
	Metadata Metadata `json:"metadata"`
}

// Metadata for a parameter.
type Metadata struct {
	// Description of a parameter.
	Description string `json:"description"`
}

// Resource represents an Azure resource in an ARM template.
type Resource struct {
	// Resource type.
	Type string `json:"type"`
	// The API Version of the resource.
	APIVersion string `json:"apiVersion"`
	// Resource name.
	Name string `json:"name"`
	// Resource location.
	Location string `json:"location"`
	// Resource tags.
	Tags map[string]interface{} `json:"tags"`
	// The SKU of the resource.
	SKU SKU `json:"sku"`
	// The kind of the resource.
	Kind string `json:"kind"`
	// The resource properties.
	Properties map[string]interface{} `json:"properties"`
	// Resource dependencies.
	DependsOn []string `json:"dependsOn"`
	// Used for specifying a scope different than the deployment scope.
	Scope string `json:"scope"`
	// Nested resources.
	Resources []Resource `json:"resources"`
}

// SKU of the namespace.
type SKU struct {
	// Name of this SKU.
	Name string `json:"name"`
	// The tier of this SKU.
	Tier string `json:"tier"`
}
