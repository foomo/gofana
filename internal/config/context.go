package config

// Context is a subset of the GrafanaConfig
// Keep in sync with https://github.com/grafana/grafanactl/blob/main/internal/config/types.go#L52
type Context struct {
	// Grafana is the Grafana configuration for this context.
	Grafana *GrafanaConfig `json:"grafana,omitempty" yaml:"grafana,omitempty"`
}
