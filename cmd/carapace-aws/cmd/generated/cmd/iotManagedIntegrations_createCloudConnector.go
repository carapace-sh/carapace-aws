package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_createCloudConnectorCmd = &cobra.Command{
	Use:   "create-cloud-connector",
	Short: "Creates a C2C (cloud-to-cloud) connector.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_createCloudConnectorCmd).Standalone()

	iotManagedIntegrations_createCloudConnectorCmd.Flags().String("client-token", "", "An idempotency token.")
	iotManagedIntegrations_createCloudConnectorCmd.Flags().String("description", "", "A description of the C2C connector.")
	iotManagedIntegrations_createCloudConnectorCmd.Flags().String("endpoint-config", "", "The configuration details for the cloud connector endpoint, including connection parameters and authentication requirements.")
	iotManagedIntegrations_createCloudConnectorCmd.Flags().String("endpoint-type", "", "The type of endpoint used for the cloud connector, which defines how the connector communicates with external services.")
	iotManagedIntegrations_createCloudConnectorCmd.Flags().String("name", "", "The display name of the C2C connector.")
	iotManagedIntegrations_createCloudConnectorCmd.MarkFlagRequired("endpoint-config")
	iotManagedIntegrations_createCloudConnectorCmd.MarkFlagRequired("name")
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_createCloudConnectorCmd)
}
