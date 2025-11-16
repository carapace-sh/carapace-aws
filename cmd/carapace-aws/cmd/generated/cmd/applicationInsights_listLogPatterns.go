package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationInsights_listLogPatternsCmd = &cobra.Command{
	Use:   "list-log-patterns",
	Short: "Lists the log patterns in the specific log `LogPatternSet`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationInsights_listLogPatternsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(applicationInsights_listLogPatternsCmd).Standalone()

		applicationInsights_listLogPatternsCmd.Flags().String("account-id", "", "The Amazon Web Services account ID for the resource group owner.")
		applicationInsights_listLogPatternsCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
		applicationInsights_listLogPatternsCmd.Flags().String("next-token", "", "The token to request the next page of results.")
		applicationInsights_listLogPatternsCmd.Flags().String("pattern-set-name", "", "The name of the log pattern set.")
		applicationInsights_listLogPatternsCmd.Flags().String("resource-group-name", "", "The name of the resource group.")
		applicationInsights_listLogPatternsCmd.MarkFlagRequired("resource-group-name")
	})
	applicationInsightsCmd.AddCommand(applicationInsights_listLogPatternsCmd)
}
