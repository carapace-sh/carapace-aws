package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudsearchCmd = &cobra.Command{
	Use:   "cloudsearch",
	Short: "Amazon CloudSearch Configuration Service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudsearchCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudsearchCmd).Standalone()

	})
	rootCmd.AddCommand(cloudsearchCmd)
}
