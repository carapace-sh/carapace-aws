package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var internetmonitor_deleteMonitorCmd = &cobra.Command{
	Use:   "delete-monitor",
	Short: "Deletes a monitor in Amazon CloudWatch Internet Monitor.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(internetmonitor_deleteMonitorCmd).Standalone()

	internetmonitor_deleteMonitorCmd.Flags().String("monitor-name", "", "The name of the monitor to delete.")
	internetmonitor_deleteMonitorCmd.MarkFlagRequired("monitor-name")
	internetmonitorCmd.AddCommand(internetmonitor_deleteMonitorCmd)
}
