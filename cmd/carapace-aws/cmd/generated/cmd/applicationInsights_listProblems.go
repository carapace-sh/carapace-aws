package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationInsights_listProblemsCmd = &cobra.Command{
	Use:   "list-problems",
	Short: "Lists the problems with your application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationInsights_listProblemsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(applicationInsights_listProblemsCmd).Standalone()

		applicationInsights_listProblemsCmd.Flags().String("account-id", "", "The Amazon Web Services account ID for the resource group owner.")
		applicationInsights_listProblemsCmd.Flags().String("component-name", "", "The name of the component.")
		applicationInsights_listProblemsCmd.Flags().String("end-time", "", "The time when the problem ended, in epoch seconds.")
		applicationInsights_listProblemsCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
		applicationInsights_listProblemsCmd.Flags().String("next-token", "", "The token to request the next page of results.")
		applicationInsights_listProblemsCmd.Flags().String("resource-group-name", "", "The name of the resource group.")
		applicationInsights_listProblemsCmd.Flags().String("start-time", "", "The time when the problem was detected, in epoch seconds.")
		applicationInsights_listProblemsCmd.Flags().String("visibility", "", "Specifies whether or not you can view the problem.")
	})
	applicationInsightsCmd.AddCommand(applicationInsights_listProblemsCmd)
}
