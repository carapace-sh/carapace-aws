package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmonitor_deleteMonitorCmd = &cobra.Command{
	Use:   "delete-monitor",
	Short: "Deletes a specified monitor.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmonitor_deleteMonitorCmd).Standalone()

	networkmonitor_deleteMonitorCmd.Flags().String("monitor-name", "", "The name of the monitor to delete.")
	networkmonitor_deleteMonitorCmd.MarkFlagRequired("monitor-name")
	networkmonitorCmd.AddCommand(networkmonitor_deleteMonitorCmd)
}
