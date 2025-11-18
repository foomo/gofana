package command

import (
	"github.com/foomo/gofana/internal/config"
	"github.com/foomo/gofana/pkg/api"
	pkgmanifest "github.com/foomo/gofana/pkg/common/manifest"
	foomokeellib "github.com/foomo/gofana/pkg/library/foomo/keel"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewDeploy() *cobra.Command {
	c := viper.New()

	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate resources",
		RunE: func(cmd *cobra.Command, args []string) error {

			cfg, err := config.Load(c, cmd)
			if err != nil {
				return err
			}

			client, err := api.NewClient(cfg.GetCurrentContext())
			if err != nil {
				return err
			}

			if err := api.FindOrCreateFolders(cmd.Context(), client, []api.Folder{
				{
					UID:  "sesamy",
					Name: "sesamy",
					Folders: []api.Folder{
						{
							UID:  "sesamy_site",
							Name: "Site",
							Folders: []api.Folder{
								{
									UID:     "sesamy_site_backend",
									Name:    "Backend",
									Folders: nil,
								},
							},
						},
					},
				},
			}, ""); err != nil {
				return err
			}

			tags := []string{"sesamy", "sesamy-site", "sesamy-site-backend"}
			builder := foomokeellib.NewServerDashboard("squadron-sesamy-site", "site-sesamy-gtm-tagging").Tags(tags)
			// builder := gotsrpc.NewServerDashboard("squadron-sesamy-site", "site-backend").Tags(tags)
			// builder := squadron.NewReleasesDashboard().Tags(tags)
			// builder := otelhttp.NewClientDashboard("squadron-sesamy-site", "site-backend").Tags(tags)

			d, err := builder.Build()
			if err != nil {
				return err
			}

			// fmt.Println(string(out))
			return pkgmanifest.Generate(pkgmanifest.Dashboard(d, "sesamy_site_backend"), c.GetBool("raw"))
			// return quick.Highlight(os.Stdout, string(out), "yaml", "terminal", "monokai")
		},
	}

	flags := cmd.Flags()

	flags.StringSliceP("config", "c", []string{"grafanactl.yaml"}, "config files (default is grafanactl.yaml)")
	_ = c.BindPFlag("config", flags.Lookup("config"))

	flags.BoolP("raw", "r", false, "Print unformatted output")
	_ = c.BindPFlag("raw", flags.Lookup("raw"))

	return cmd
}
