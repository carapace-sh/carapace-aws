package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_updateDestinationCmd = &cobra.Command{
	Use:   "update-destination",
	Short: "Update a destination specified by name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_updateDestinationCmd).Standalone()

	iotManagedIntegrations_updateDestinationCmd.Flags().String("delivery-destination-arn", "", "The Amazon Resource Name (ARN) of the customer-managed destination.")
	iotManagedIntegrations_updateDestinationCmd.Flags().String("delivery-destination-type", "", "The destination type for the customer-managed destination.")
	iotManagedIntegrations_updateDestinationCmd.Flags().String("description", "", "The description of the customer-managed destination.")
	iotManagedIntegrations_updateDestinationCmd.Flags().String("name", "", "The name of the customer-managed destination.")
	iotManagedIntegrations_updateDestinationCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of the delivery destination role.")
	iotManagedIntegrations_updateDestinationCmd.MarkFlagRequired("name")
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_updateDestinationCmd)
}
