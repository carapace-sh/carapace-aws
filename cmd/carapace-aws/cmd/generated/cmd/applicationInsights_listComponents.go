package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationInsights_listComponentsCmd = &cobra.Command{
	Use:   "list-components",
	Short: "Lists the auto-grouped, standalone, and custom components of the application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationInsights_listComponentsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(applicationInsights_listComponentsCmd).Standalone()

		applicationInsights_listComponentsCmd.Flags().String("account-id", "", "The Amazon Web Services account ID for the resource group owner.")
		applicationInsights_listComponentsCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
		applicationInsights_listComponentsCmd.Flags().String("next-token", "", "The token to request the next page of results.")
		applicationInsights_listComponentsCmd.Flags().String("resource-group-name", "", "The name of the resource group.")
		applicationInsights_listComponentsCmd.MarkFlagRequired("resource-group-name")
	})
	applicationInsightsCmd.AddCommand(applicationInsights_listComponentsCmd)
}
