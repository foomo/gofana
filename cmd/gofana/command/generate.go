package command

import (
	"os"
	"os/exec"

	"github.com/foomo/gofana/internal/gofana"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewGenerate() *cobra.Command {
	c := viper.New()

	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate resources",
		RunE: func(cmd *cobra.Command, args []string) error {
			paths, err := gofana.List(c.GetStringSlice("path"))
			if err != nil {
				return err
			}

			for _, path := range paths {
				sh := exec.CommandContext(cmd.Context(), "go", "run", path, "generate")
				if c.GetBool("raw") {
					sh.Args = append(sh.Args, "--raw")
				}

				sh.Env = os.Environ()
				sh.Stdout = os.Stdout
				sh.Stderr = os.Stderr
				pterm.Debug.Println("running command:", sh.String())

				if err := sh.Run(); err != nil {
					return err
				}
			}

			return nil
		},
	}

	flags := cmd.Flags()

	flags.StringArray("path", []string{}, "path to search for gofana resources")
	_ = c.BindPFlag("path", flags.Lookup("path"))

	flags.BoolP("raw", "r", false, "Print unformatted output")
	_ = c.BindPFlag("raw", flags.Lookup("raw"))

	return cmd
}
