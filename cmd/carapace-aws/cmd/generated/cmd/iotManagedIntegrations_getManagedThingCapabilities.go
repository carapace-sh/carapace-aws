package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_getManagedThingCapabilitiesCmd = &cobra.Command{
	Use:   "get-managed-thing-capabilities",
	Short: "Get the capabilities for a managed thing using the device ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_getManagedThingCapabilitiesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotManagedIntegrations_getManagedThingCapabilitiesCmd).Standalone()

		iotManagedIntegrations_getManagedThingCapabilitiesCmd.Flags().String("identifier", "", "The id of the device.")
		iotManagedIntegrations_getManagedThingCapabilitiesCmd.MarkFlagRequired("identifier")
	})
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_getManagedThingCapabilitiesCmd)
}
