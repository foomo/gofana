package plugin

import (
	"github.com/foomo/gofana/pkg/api"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// NewRoot represents the base command when called without any subcommands
func NewRoot(folderUID string, folder api.Folder, resources []any) *cobra.Command {
	c := viper.New()

	cmd := &cobra.Command{
		Use:           "gofana",
		Short:         "Gofana plugin",
		SilenceErrors: true,
		SilenceUsage:  true,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			pterm.PrintDebugMessages = c.GetBool("debug")
		},
	}

	flags := cmd.PersistentFlags()

	flags.StringSliceP("config", "c", []string{"grafanactl.yaml"}, "config files (default is grafana.yaml)")
	_ = c.BindPFlag("config", flags.Lookup("config"))

	flags.BoolP("debug", "v", false, "output debug information")
	_ = c.BindPFlag("debug", flags.Lookup("debug"))

	cmd.AddCommand(
		NewConfig(),
		NewGenerate(folderUID, resources),
		NewDeploy(folderUID, folder, resources),
	)

	return cmd
}
