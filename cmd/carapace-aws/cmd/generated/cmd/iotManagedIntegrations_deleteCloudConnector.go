package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_deleteCloudConnectorCmd = &cobra.Command{
	Use:   "delete-cloud-connector",
	Short: "Delete a cloud connector.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_deleteCloudConnectorCmd).Standalone()

	iotManagedIntegrations_deleteCloudConnectorCmd.Flags().String("identifier", "", "The identifier of the cloud connector.")
	iotManagedIntegrations_deleteCloudConnectorCmd.MarkFlagRequired("identifier")
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_deleteCloudConnectorCmd)
}
