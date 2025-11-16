package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var costOptimizationHub_getPreferencesCmd = &cobra.Command{
	Use:   "get-preferences",
	Short: "Returns a set of preferences for an account in order to add account-specific preferences into the service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(costOptimizationHub_getPreferencesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(costOptimizationHub_getPreferencesCmd).Standalone()

	})
	costOptimizationHubCmd.AddCommand(costOptimizationHub_getPreferencesCmd)
}
