package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_updateCloudConnectorCmd = &cobra.Command{
	Use:   "update-cloud-connector",
	Short: "Update an existing cloud connector.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_updateCloudConnectorCmd).Standalone()

	iotManagedIntegrations_updateCloudConnectorCmd.Flags().String("description", "", "The new description to assign to the cloud connector.")
	iotManagedIntegrations_updateCloudConnectorCmd.Flags().String("identifier", "", "The unique identifier of the cloud connector to update.")
	iotManagedIntegrations_updateCloudConnectorCmd.Flags().String("name", "", "The new display name to assign to the cloud connector.")
	iotManagedIntegrations_updateCloudConnectorCmd.MarkFlagRequired("identifier")
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_updateCloudConnectorCmd)
}
