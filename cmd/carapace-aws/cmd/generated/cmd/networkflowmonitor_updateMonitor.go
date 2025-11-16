package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkflowmonitor_updateMonitorCmd = &cobra.Command{
	Use:   "update-monitor",
	Short: "Update a monitor to add or remove local or remote resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkflowmonitor_updateMonitorCmd).Standalone()

	networkflowmonitor_updateMonitorCmd.Flags().String("client-token", "", "A unique, case-sensitive string of up to 64 ASCII characters that you specify to make an idempotent API request.")
	networkflowmonitor_updateMonitorCmd.Flags().String("local-resources-to-add", "", "Additional local resources to specify network flows for a monitor, as an array of resources with identifiers and types.")
	networkflowmonitor_updateMonitorCmd.Flags().String("local-resources-to-remove", "", "The local resources to remove, as an array of resources with identifiers and types.")
	networkflowmonitor_updateMonitorCmd.Flags().String("monitor-name", "", "The name of the monitor.")
	networkflowmonitor_updateMonitorCmd.Flags().String("remote-resources-to-add", "", "The remote resources to add, as an array of resources with identifiers and types.")
	networkflowmonitor_updateMonitorCmd.Flags().String("remote-resources-to-remove", "", "The remote resources to remove, as an array of resources with identifiers and types.")
	networkflowmonitor_updateMonitorCmd.MarkFlagRequired("monitor-name")
	networkflowmonitorCmd.AddCommand(networkflowmonitor_updateMonitorCmd)
}
