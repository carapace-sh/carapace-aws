package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_createDestinationCmd = &cobra.Command{
	Use:   "create-destination",
	Short: "Create a notification destination such as Kinesis Data Streams that receive events and notifications from Managed integrations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_createDestinationCmd).Standalone()

	iotManagedIntegrations_createDestinationCmd.Flags().String("client-token", "", "An idempotency token.")
	iotManagedIntegrations_createDestinationCmd.Flags().String("delivery-destination-arn", "", "The Amazon Resource Name (ARN) of the customer-managed destination.")
	iotManagedIntegrations_createDestinationCmd.Flags().String("delivery-destination-type", "", "The destination type for the customer-managed destination.")
	iotManagedIntegrations_createDestinationCmd.Flags().String("description", "", "The description of the customer-managed destination.")
	iotManagedIntegrations_createDestinationCmd.Flags().String("name", "", "The name of the customer-managed destination.")
	iotManagedIntegrations_createDestinationCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of the delivery destination role.")
	iotManagedIntegrations_createDestinationCmd.Flags().String("tags", "", "A set of key/value pairs that are used to manage the destination.")
	iotManagedIntegrations_createDestinationCmd.MarkFlagRequired("delivery-destination-arn")
	iotManagedIntegrations_createDestinationCmd.MarkFlagRequired("delivery-destination-type")
	iotManagedIntegrations_createDestinationCmd.MarkFlagRequired("name")
	iotManagedIntegrations_createDestinationCmd.MarkFlagRequired("role-arn")
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_createDestinationCmd)
}
