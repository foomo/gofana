package command

import (
	"strings"

	"github.com/foomo/gofana/internal/gofana"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewList() *cobra.Command {
	c := viper.New()

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List gofana resources",
		RunE: func(cmd *cobra.Command, args []string) error {

			pterm.Debug.Println("Searching for gofana resources...", c.GetStringSlice("path"))
			paths, err := gofana.List(c.GetStringSlice("path"))
			if err != nil {
				return err
			}

			pterm.Info.Println(strings.Join(paths, "\n"))
			return nil
		},
	}

	flags := cmd.Flags()

	flags.StringArray("path", []string{}, "path to search for gofana resources")
	_ = c.BindPFlag("path", flags.Lookup("path"))

	return cmd
}
