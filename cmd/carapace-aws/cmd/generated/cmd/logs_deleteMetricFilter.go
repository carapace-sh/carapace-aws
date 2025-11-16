package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_deleteMetricFilterCmd = &cobra.Command{
	Use:   "delete-metric-filter",
	Short: "Deletes the specified metric filter.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_deleteMetricFilterCmd).Standalone()

	logs_deleteMetricFilterCmd.Flags().String("filter-name", "", "The name of the metric filter.")
	logs_deleteMetricFilterCmd.Flags().String("log-group-name", "", "The name of the log group.")
	logs_deleteMetricFilterCmd.MarkFlagRequired("filter-name")
	logs_deleteMetricFilterCmd.MarkFlagRequired("log-group-name")
	logsCmd.AddCommand(logs_deleteMetricFilterCmd)
}
