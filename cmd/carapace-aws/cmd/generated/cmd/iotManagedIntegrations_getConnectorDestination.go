package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_getConnectorDestinationCmd = &cobra.Command{
	Use:   "get-connector-destination",
	Short: "Get connector destination details linked to a cloud-to-cloud (C2C) connector.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_getConnectorDestinationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotManagedIntegrations_getConnectorDestinationCmd).Standalone()

		iotManagedIntegrations_getConnectorDestinationCmd.Flags().String("identifier", "", "The identifier of the C2C connector destination.")
		iotManagedIntegrations_getConnectorDestinationCmd.MarkFlagRequired("identifier")
	})
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_getConnectorDestinationCmd)
}
