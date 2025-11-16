package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var internetmonitor_stopQueryCmd = &cobra.Command{
	Use:   "stop-query",
	Short: "Stop a query that is progress for a specific monitor.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(internetmonitor_stopQueryCmd).Standalone()

	internetmonitor_stopQueryCmd.Flags().String("monitor-name", "", "The name of the monitor.")
	internetmonitor_stopQueryCmd.Flags().String("query-id", "", "The ID of the query that you want to stop.")
	internetmonitor_stopQueryCmd.MarkFlagRequired("monitor-name")
	internetmonitor_stopQueryCmd.MarkFlagRequired("query-id")
	internetmonitorCmd.AddCommand(internetmonitor_stopQueryCmd)
}
