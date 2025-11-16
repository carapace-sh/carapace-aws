package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_getManagedThingConnectivityDataCmd = &cobra.Command{
	Use:   "get-managed-thing-connectivity-data",
	Short: "Get the connectivity status of a managed thing.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_getManagedThingConnectivityDataCmd).Standalone()

	iotManagedIntegrations_getManagedThingConnectivityDataCmd.Flags().String("identifier", "", "The identifier of a managed thing.")
	iotManagedIntegrations_getManagedThingConnectivityDataCmd.MarkFlagRequired("identifier")
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_getManagedThingConnectivityDataCmd)
}
