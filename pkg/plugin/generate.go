package plugin

import (
	pkgmanifest "github.com/foomo/gofana/pkg/common/manifest"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewGenerate(folderUID string, resources []any) *cobra.Command {
	c := viper.New()

	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate resources",
		RunE: func(cmd *cobra.Command, args []string) error {
			for _, resource := range resources {
				switch t := resource.(type) {
				case *dashboard.DashboardBuilder:
					d, err := t.Build()
					if err != nil {
						return err
					}

					if err := pkgmanifest.Generate(pkgmanifest.Dashboard(d, folderUID), c.GetBool("raw")); err != nil {
						return err
					}
				default:
					pterm.Fatal.Printfln("unknown resource type: %v", resource)
				}
			}

			return nil
		},
	}

	flags := cmd.Flags()

	flags.BoolP("raw", "r", false, "Print unformatted output")
	_ = c.BindPFlag("raw", flags.Lookup("raw"))

	return cmd
}
