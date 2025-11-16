package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_testMetricFilterCmd = &cobra.Command{
	Use:   "test-metric-filter",
	Short: "Tests the filter pattern of a metric filter against a sample of log event messages.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_testMetricFilterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(logs_testMetricFilterCmd).Standalone()

		logs_testMetricFilterCmd.Flags().String("filter-pattern", "", "")
		logs_testMetricFilterCmd.Flags().String("log-event-messages", "", "The log event messages to test.")
		logs_testMetricFilterCmd.MarkFlagRequired("filter-pattern")
		logs_testMetricFilterCmd.MarkFlagRequired("log-event-messages")
	})
	logsCmd.AddCommand(logs_testMetricFilterCmd)
}
