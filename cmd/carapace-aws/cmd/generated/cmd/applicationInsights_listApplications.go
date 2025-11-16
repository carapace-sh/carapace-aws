package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationInsights_listApplicationsCmd = &cobra.Command{
	Use:   "list-applications",
	Short: "Lists the IDs of the applications that you are monitoring.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationInsights_listApplicationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(applicationInsights_listApplicationsCmd).Standalone()

		applicationInsights_listApplicationsCmd.Flags().String("account-id", "", "The Amazon Web Services account ID for the resource group owner.")
		applicationInsights_listApplicationsCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
		applicationInsights_listApplicationsCmd.Flags().String("next-token", "", "The token to request the next page of results.")
	})
	applicationInsightsCmd.AddCommand(applicationInsights_listApplicationsCmd)
}
