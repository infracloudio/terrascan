package config

// config holds the common resource config fields
type config struct {
	Location string            `json:"location"`
	Name     string            `json:"name"`
	Tags     map[string]string `json:"tags"`
}
