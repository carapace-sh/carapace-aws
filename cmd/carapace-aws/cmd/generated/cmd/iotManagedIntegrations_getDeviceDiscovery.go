package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_getDeviceDiscoveryCmd = &cobra.Command{
	Use:   "get-device-discovery",
	Short: "Get the current state of a device discovery.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_getDeviceDiscoveryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotManagedIntegrations_getDeviceDiscoveryCmd).Standalone()

		iotManagedIntegrations_getDeviceDiscoveryCmd.Flags().String("identifier", "", "The id of the device discovery job request.")
		iotManagedIntegrations_getDeviceDiscoveryCmd.MarkFlagRequired("identifier")
	})
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_getDeviceDiscoveryCmd)
}
