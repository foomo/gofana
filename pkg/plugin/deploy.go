package plugin

import (
	"github.com/foomo/gofana/internal/config"
	"github.com/foomo/gofana/pkg/api"
	pkgmanifest "github.com/foomo/gofana/pkg/common/manifest"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewDeploy(folderUID string, folder api.Folder, resources []any) *cobra.Command {
	c := viper.New()

	cmd := &cobra.Command{
		Use:   "deploy",
		Short: "Generate & deploy resources",
		RunE: func(cmd *cobra.Command, args []string) error {

			cfg, err := config.Load(c, cmd)
			if err != nil {
				return err
			}

			client, err := api.NewClient(cfg.GetCurrentContext())
			if err != nil {
				return err
			}

			if err := api.FindOrCreateFolders(cmd.Context(), client, []api.Folder{folder}, ""); err != nil {
				return err
			}

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

	flags.StringSliceP("config", "c", []string{"grafanactl.yaml"}, "config files (default is grafanactl.yaml)")
	_ = c.BindPFlag("config", flags.Lookup("config"))

	return cmd
}
