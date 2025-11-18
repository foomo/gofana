package api

import (
	"net/url"
	"strings"

	"github.com/foomo/gofana/internal/config"
	"github.com/go-openapi/strfmt"
	"github.com/grafana/grafana-openapi-client-go/client"
	"github.com/pkg/errors"
)

func NewClient(ctx *config.Context) (*client.GrafanaHTTPAPI, error) {
	if ctx == nil {
		return nil, errors.New("no context provided")
	}
	if ctx.Grafana == nil {
		return nil, errors.New("grafana not configured")
	}

	grafanaURL, err := url.Parse(ctx.Grafana.Server)
	if err != nil {
		return nil, err
	}

	cfg := &client.TransportConfig{
		Host:     grafanaURL.Host,
		BasePath: strings.TrimLeft(grafanaURL.Path+"/api", "/"),
		Schemes:  []string{grafanaURL.Scheme},
	}

	if ctx.Grafana.TLS != nil {
		cfg.TLSConfig = ctx.Grafana.TLS.ToStdTLSConfig()
	}

	// Authentication
	if ctx.Grafana.User != "" && ctx.Grafana.Password != "" {
		cfg.BasicAuth = url.UserPassword(ctx.Grafana.User, ctx.Grafana.Password)
	}
	if ctx.Grafana.APIToken != "" {
		cfg.APIKey = ctx.Grafana.APIToken
	}
	if ctx.Grafana.OrgID != 0 {
		cfg.OrgID = ctx.Grafana.OrgID
	}

	return client.NewHTTPClientWithConfig(strfmt.Default, cfg), nil
}
