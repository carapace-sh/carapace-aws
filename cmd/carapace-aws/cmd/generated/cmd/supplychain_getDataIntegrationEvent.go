package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var supplychain_getDataIntegrationEventCmd = &cobra.Command{
	Use:   "get-data-integration-event",
	Short: "Enables you to programmatically view an Amazon Web Services Supply Chain Data Integration Event.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(supplychain_getDataIntegrationEventCmd).Standalone()

	supplychain_getDataIntegrationEventCmd.Flags().String("event-id", "", "The unique event identifier.")
	supplychain_getDataIntegrationEventCmd.Flags().String("instance-id", "", "The Amazon Web Services Supply Chain instance identifier.")
	supplychain_getDataIntegrationEventCmd.MarkFlagRequired("event-id")
	supplychain_getDataIntegrationEventCmd.MarkFlagRequired("instance-id")
	supplychainCmd.AddCommand(supplychain_getDataIntegrationEventCmd)
}
