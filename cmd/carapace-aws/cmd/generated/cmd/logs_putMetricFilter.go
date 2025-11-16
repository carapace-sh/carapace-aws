package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_putMetricFilterCmd = &cobra.Command{
	Use:   "put-metric-filter",
	Short: "Creates or updates a metric filter and associates it with the specified log group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_putMetricFilterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(logs_putMetricFilterCmd).Standalone()

		logs_putMetricFilterCmd.Flags().String("apply-on-transformed-logs", "", "This parameter is valid only for log groups that have an active log transformer.")
		logs_putMetricFilterCmd.Flags().String("emit-system-field-dimensions", "", "A list of system fields to emit as additional dimensions in the generated metrics.")
		logs_putMetricFilterCmd.Flags().String("field-selection-criteria", "", "A filter expression that specifies which log events should be processed by this metric filter based on system fields such as source account and source region.")
		logs_putMetricFilterCmd.Flags().String("filter-name", "", "A name for the metric filter.")
		logs_putMetricFilterCmd.Flags().String("filter-pattern", "", "A filter pattern for extracting metric data out of ingested log events.")
		logs_putMetricFilterCmd.Flags().String("log-group-name", "", "The name of the log group.")
		logs_putMetricFilterCmd.Flags().String("metric-transformations", "", "A collection of information that defines how metric data gets emitted.")
		logs_putMetricFilterCmd.MarkFlagRequired("filter-name")
		logs_putMetricFilterCmd.MarkFlagRequired("filter-pattern")
		logs_putMetricFilterCmd.MarkFlagRequired("log-group-name")
		logs_putMetricFilterCmd.MarkFlagRequired("metric-transformations")
	})
	logsCmd.AddCommand(logs_putMetricFilterCmd)
}
