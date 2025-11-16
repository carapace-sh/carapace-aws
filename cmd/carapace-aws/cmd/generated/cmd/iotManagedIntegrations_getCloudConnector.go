package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_getCloudConnectorCmd = &cobra.Command{
	Use:   "get-cloud-connector",
	Short: "Get configuration details for a cloud connector.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_getCloudConnectorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotManagedIntegrations_getCloudConnectorCmd).Standalone()

		iotManagedIntegrations_getCloudConnectorCmd.Flags().String("identifier", "", "The identifier of the C2C connector.")
		iotManagedIntegrations_getCloudConnectorCmd.MarkFlagRequired("identifier")
	})
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_getCloudConnectorCmd)
}
