package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var supplychain_listDataIntegrationEventsCmd = &cobra.Command{
	Use:   "list-data-integration-events",
	Short: "Enables you to programmatically list all data integration events for the provided Amazon Web Services Supply Chain instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(supplychain_listDataIntegrationEventsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(supplychain_listDataIntegrationEventsCmd).Standalone()

		supplychain_listDataIntegrationEventsCmd.Flags().String("event-type", "", "List data integration events for the specified eventType.")
		supplychain_listDataIntegrationEventsCmd.Flags().String("instance-id", "", "The Amazon Web Services Supply Chain instance identifier.")
		supplychain_listDataIntegrationEventsCmd.Flags().String("max-results", "", "Specify the maximum number of data integration events to fetch in one paginated request.")
		supplychain_listDataIntegrationEventsCmd.Flags().String("next-token", "", "The pagination token to fetch the next page of the data integration events.")
		supplychain_listDataIntegrationEventsCmd.MarkFlagRequired("instance-id")
	})
	supplychainCmd.AddCommand(supplychain_listDataIntegrationEventsCmd)
}
