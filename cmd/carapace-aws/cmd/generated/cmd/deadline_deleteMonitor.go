package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_deleteMonitorCmd = &cobra.Command{
	Use:   "delete-monitor",
	Short: "Removes a Deadline Cloud monitor.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_deleteMonitorCmd).Standalone()

	deadline_deleteMonitorCmd.Flags().String("monitor-id", "", "The unique identifier of the monitor to delete.")
	deadline_deleteMonitorCmd.MarkFlagRequired("monitor-id")
	deadlineCmd.AddCommand(deadline_deleteMonitorCmd)
}
