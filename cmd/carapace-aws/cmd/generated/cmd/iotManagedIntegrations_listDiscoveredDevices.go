package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_listDiscoveredDevicesCmd = &cobra.Command{
	Use:   "list-discovered-devices",
	Short: "Lists all devices discovered during a specific device discovery task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_listDiscoveredDevicesCmd).Standalone()

	iotManagedIntegrations_listDiscoveredDevicesCmd.Flags().String("identifier", "", "The identifier of the device discovery job to list discovered devices for.")
	iotManagedIntegrations_listDiscoveredDevicesCmd.Flags().String("max-results", "", "The maximum number of discovered devices to return in a single response.")
	iotManagedIntegrations_listDiscoveredDevicesCmd.Flags().String("next-token", "", "A token used for pagination of results.")
	iotManagedIntegrations_listDiscoveredDevicesCmd.MarkFlagRequired("identifier")
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_listDiscoveredDevicesCmd)
}
