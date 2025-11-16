package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_updateConnectorDestinationCmd = &cobra.Command{
	Use:   "update-connector-destination",
	Short: "Updates the properties of an existing connector destination.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_updateConnectorDestinationCmd).Standalone()

	iotManagedIntegrations_updateConnectorDestinationCmd.Flags().String("auth-config", "", "The updated authentication configuration details for the connector destination.")
	iotManagedIntegrations_updateConnectorDestinationCmd.Flags().String("auth-type", "", "The new authentication type to use for the connector destination.")
	iotManagedIntegrations_updateConnectorDestinationCmd.Flags().String("description", "", "The new description to assign to the connector destination.")
	iotManagedIntegrations_updateConnectorDestinationCmd.Flags().String("identifier", "", "The unique identifier of the connector destination to update.")
	iotManagedIntegrations_updateConnectorDestinationCmd.Flags().String("name", "", "The new display name to assign to the connector destination.")
	iotManagedIntegrations_updateConnectorDestinationCmd.Flags().String("secrets-manager", "", "The updated AWS Secrets Manager configuration for the connector destination.")
	iotManagedIntegrations_updateConnectorDestinationCmd.MarkFlagRequired("identifier")
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_updateConnectorDestinationCmd)
}
