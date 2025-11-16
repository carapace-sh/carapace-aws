package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkflowmonitor_createMonitorCmd = &cobra.Command{
	Use:   "create-monitor",
	Short: "Create a monitor for specific network flows between local and remote resources, so that you can monitor network performance for one or several of your workloads.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkflowmonitor_createMonitorCmd).Standalone()

	networkflowmonitor_createMonitorCmd.Flags().String("client-token", "", "A unique, case-sensitive string of up to 64 ASCII characters that you specify to make an idempotent API request.")
	networkflowmonitor_createMonitorCmd.Flags().String("local-resources", "", "The local resources to monitor.")
	networkflowmonitor_createMonitorCmd.Flags().String("monitor-name", "", "The name of the monitor.")
	networkflowmonitor_createMonitorCmd.Flags().String("remote-resources", "", "The remote resources to monitor.")
	networkflowmonitor_createMonitorCmd.Flags().String("scope-arn", "", "The Amazon Resource Name (ARN) of the scope for the monitor.")
	networkflowmonitor_createMonitorCmd.Flags().String("tags", "", "The tags for a monitor.")
	networkflowmonitor_createMonitorCmd.MarkFlagRequired("local-resources")
	networkflowmonitor_createMonitorCmd.MarkFlagRequired("monitor-name")
	networkflowmonitor_createMonitorCmd.MarkFlagRequired("scope-arn")
	networkflowmonitorCmd.AddCommand(networkflowmonitor_createMonitorCmd)
}
