package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_deleteConnectorDestinationCmd = &cobra.Command{
	Use:   "delete-connector-destination",
	Short: "Delete a connector destination linked to a cloud-to-cloud (C2C) connector.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_deleteConnectorDestinationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotManagedIntegrations_deleteConnectorDestinationCmd).Standalone()

		iotManagedIntegrations_deleteConnectorDestinationCmd.Flags().String("identifier", "", "The identifier of the connector destination.")
		iotManagedIntegrations_deleteConnectorDestinationCmd.MarkFlagRequired("identifier")
	})
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_deleteConnectorDestinationCmd)
}
