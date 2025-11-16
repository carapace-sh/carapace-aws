package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationInsights_listConfigurationHistoryCmd = &cobra.Command{
	Use:   "list-configuration-history",
	Short: "Lists the INFO, WARN, and ERROR events for periodic configuration updates performed by Application Insights.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationInsights_listConfigurationHistoryCmd).Standalone()

	applicationInsights_listConfigurationHistoryCmd.Flags().String("account-id", "", "The Amazon Web Services account ID for the resource group owner.")
	applicationInsights_listConfigurationHistoryCmd.Flags().String("end-time", "", "The end time of the event.")
	applicationInsights_listConfigurationHistoryCmd.Flags().String("event-status", "", "The status of the configuration update event.")
	applicationInsights_listConfigurationHistoryCmd.Flags().String("max-results", "", "The maximum number of results returned by `ListConfigurationHistory` in paginated output.")
	applicationInsights_listConfigurationHistoryCmd.Flags().String("next-token", "", "The `NextToken` value returned from a previous paginated `ListConfigurationHistory` request where `MaxResults` was used and the results exceeded the value of that parameter.")
	applicationInsights_listConfigurationHistoryCmd.Flags().String("resource-group-name", "", "Resource group to which the application belongs.")
	applicationInsights_listConfigurationHistoryCmd.Flags().String("start-time", "", "The start time of the event.")
	applicationInsightsCmd.AddCommand(applicationInsights_listConfigurationHistoryCmd)
}
