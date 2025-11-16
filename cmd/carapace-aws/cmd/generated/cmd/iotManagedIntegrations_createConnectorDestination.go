package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_createConnectorDestinationCmd = &cobra.Command{
	Use:   "create-connector-destination",
	Short: "Create a connector destination for connecting a cloud-to-cloud (C2C) connector to the customer's Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_createConnectorDestinationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotManagedIntegrations_createConnectorDestinationCmd).Standalone()

		iotManagedIntegrations_createConnectorDestinationCmd.Flags().String("auth-config", "", "The authentication configuration details for the connector destination, including OAuth settings and other authentication parameters.")
		iotManagedIntegrations_createConnectorDestinationCmd.Flags().String("auth-type", "", "The authentication type used for the connector destination, which determines how credentials and access are managed.")
		iotManagedIntegrations_createConnectorDestinationCmd.Flags().String("client-token", "", "An idempotency token.")
		iotManagedIntegrations_createConnectorDestinationCmd.Flags().String("cloud-connector-id", "", "The identifier of the C2C connector.")
		iotManagedIntegrations_createConnectorDestinationCmd.Flags().String("description", "", "A description of the connector destination.")
		iotManagedIntegrations_createConnectorDestinationCmd.Flags().String("name", "", "The display name of the connector destination.")
		iotManagedIntegrations_createConnectorDestinationCmd.Flags().String("secrets-manager", "", "The AWS Secrets Manager configuration used to securely store and manage sensitive information for the connector destination.")
		iotManagedIntegrations_createConnectorDestinationCmd.MarkFlagRequired("auth-config")
		iotManagedIntegrations_createConnectorDestinationCmd.MarkFlagRequired("auth-type")
		iotManagedIntegrations_createConnectorDestinationCmd.MarkFlagRequired("cloud-connector-id")
		iotManagedIntegrations_createConnectorDestinationCmd.MarkFlagRequired("secrets-manager")
	})
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_createConnectorDestinationCmd)
}
