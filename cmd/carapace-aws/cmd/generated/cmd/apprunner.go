package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apprunnerCmd = &cobra.Command{
	Use:   "apprunner",
	Short: "App Runner",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apprunnerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apprunnerCmd).Standalone()

	})
	rootCmd.AddCommand(apprunnerCmd)
}
