package variable

import (
	dashboard2 "github.com/foomo/gofana/pkg/common/dashboard"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

func NewMultiQueryVariable(label, expr string) *dashboard.QueryVariableBuilder {
	return NewQueryVariable(label, expr).
		Multi(true).
		IncludeAll(true)
}

func NewQueryVariable(label, expr string) *dashboard.QueryVariableBuilder {
	return dashboard.NewQueryVariableBuilder(label).
		Query(dashboard2.StringOrMapAsString("label_values(" + expr + ", " + label + ")")).
		AllowCustomValue(false).
		Refresh(dashboard.VariableRefreshOnTimeRangeChanged).
		Sort(dashboard.VariableSortAlphabeticalAsc)
}

func NewNamespaceQuery(metric string) *dashboard.QueryVariableBuilder {
	return NewMultiQueryVariable("namespace", metric).Label("📂︎")
}

func NewServiceQuery(expr string) *dashboard.QueryVariableBuilder {
	return NewQueryVariable("service", expr).Label("🛎️")
}
