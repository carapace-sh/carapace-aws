package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var xrayCmd = &cobra.Command{
	Use:   "xray",
	Short: "Amazon Web Services X-Ray provides APIs for managing debug traces and retrieving service maps and other data created by processing those traces.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(xrayCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(xrayCmd).Standalone()

	})
	rootCmd.AddCommand(xrayCmd)
}
