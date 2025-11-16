package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_describeMetricFiltersCmd = &cobra.Command{
	Use:   "describe-metric-filters",
	Short: "Lists the specified metric filters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_describeMetricFiltersCmd).Standalone()

	logs_describeMetricFiltersCmd.Flags().String("filter-name-prefix", "", "The prefix to match.")
	logs_describeMetricFiltersCmd.Flags().String("limit", "", "The maximum number of items returned.")
	logs_describeMetricFiltersCmd.Flags().String("log-group-name", "", "The name of the log group.")
	logs_describeMetricFiltersCmd.Flags().String("metric-name", "", "Filters results to include only those with the specified metric name.")
	logs_describeMetricFiltersCmd.Flags().String("metric-namespace", "", "Filters results to include only those in the specified namespace.")
	logs_describeMetricFiltersCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
	logsCmd.AddCommand(logs_describeMetricFiltersCmd)
}
