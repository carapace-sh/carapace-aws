package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_getMonitorCmd = &cobra.Command{
	Use:   "get-monitor",
	Short: "Gets information about the specified monitor.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_getMonitorCmd).Standalone()

	deadline_getMonitorCmd.Flags().String("monitor-id", "", "The unique identifier for the monitor.")
	deadline_getMonitorCmd.MarkFlagRequired("monitor-id")
	deadlineCmd.AddCommand(deadline_getMonitorCmd)
}
