package variable

import (
	dashboard2 "github.com/foomo/gofana/pkg/common/dashboard"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

func ResolutionVariable() *dashboard.IntervalVariableBuilder {
	return dashboard.NewIntervalVariableBuilder("resolution").
		Label("👓").
		AllowCustomValue(false).
		Values(dashboard2.StringOrMapAsString("0.5,1,2,4")).
		Current(dashboard.VariableOption{Value: dashboard2.StringOrArrayOfStringAsString("1")})
}
