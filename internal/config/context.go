package config

// https://github.com/grafana/grafanactl/blob/main/internal/config/types.go#L52
type Context struct {
	Name    string         `json:"-" yaml:"-"`
	Grafana *GrafanaConfig `json:"grafana,omitempty" yaml:"grafana,omitempty"`
}
