package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_listDeviceDiscoveriesCmd = &cobra.Command{
	Use:   "list-device-discoveries",
	Short: "Lists all device discovery tasks, with optional filtering by type and status.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_listDeviceDiscoveriesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotManagedIntegrations_listDeviceDiscoveriesCmd).Standalone()

		iotManagedIntegrations_listDeviceDiscoveriesCmd.Flags().String("max-results", "", "The maximum number of device discovery jobs to return in a single response.")
		iotManagedIntegrations_listDeviceDiscoveriesCmd.Flags().String("next-token", "", "A token used for pagination of results.")
		iotManagedIntegrations_listDeviceDiscoveriesCmd.Flags().String("status-filter", "", "The status to filter device discovery jobs by.")
		iotManagedIntegrations_listDeviceDiscoveriesCmd.Flags().String("type-filter", "", "The discovery type to filter device discovery jobs by.")
	})
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_listDeviceDiscoveriesCmd)
}
