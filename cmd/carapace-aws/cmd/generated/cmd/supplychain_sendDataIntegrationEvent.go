package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var supplychain_sendDataIntegrationEventCmd = &cobra.Command{
	Use:   "send-data-integration-event",
	Short: "Send the data payload for the event with real-time data for analysis or monitoring.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(supplychain_sendDataIntegrationEventCmd).Standalone()

	supplychain_sendDataIntegrationEventCmd.Flags().String("client-token", "", "The idempotent client token.")
	supplychain_sendDataIntegrationEventCmd.Flags().String("data", "", "The data payload of the event, should follow the data schema of the target dataset, or see [Data entities supported in AWS Supply Chain](https://docs.aws.amazon.com/aws-supply-chain/latest/userguide/data-model-asc.html). To send single data record, use JsonObject format; to send multiple data records, use JsonArray format.")
	supplychain_sendDataIntegrationEventCmd.Flags().String("dataset-target", "", "The target dataset configuration for **scn.data.dataset** event type.")
	supplychain_sendDataIntegrationEventCmd.Flags().String("event-group-id", "", "Event identifier (for example, orderId for InboundOrder) used for data sharding or partitioning.")
	supplychain_sendDataIntegrationEventCmd.Flags().String("event-timestamp", "", "The timestamp (in epoch seconds) associated with the event.")
	supplychain_sendDataIntegrationEventCmd.Flags().String("event-type", "", "The data event type.")
	supplychain_sendDataIntegrationEventCmd.Flags().String("instance-id", "", "The AWS Supply Chain instance identifier.")
	supplychain_sendDataIntegrationEventCmd.MarkFlagRequired("data")
	supplychain_sendDataIntegrationEventCmd.MarkFlagRequired("event-group-id")
	supplychain_sendDataIntegrationEventCmd.MarkFlagRequired("event-type")
	supplychain_sendDataIntegrationEventCmd.MarkFlagRequired("instance-id")
	supplychainCmd.AddCommand(supplychain_sendDataIntegrationEventCmd)
}
