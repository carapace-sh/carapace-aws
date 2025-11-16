package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudwatch_getInsightRuleReportCmd = &cobra.Command{
	Use:   "get-insight-rule-report",
	Short: "This operation returns the time series data collected by a Contributor Insights rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudwatch_getInsightRuleReportCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudwatch_getInsightRuleReportCmd).Standalone()

		cloudwatch_getInsightRuleReportCmd.Flags().String("end-time", "", "The end time of the data to use in the report.")
		cloudwatch_getInsightRuleReportCmd.Flags().String("max-contributor-count", "", "The maximum number of contributors to include in the report.")
		cloudwatch_getInsightRuleReportCmd.Flags().String("metrics", "", "Specifies which metrics to use for aggregation of contributor values for the report.")
		cloudwatch_getInsightRuleReportCmd.Flags().String("order-by", "", "Determines what statistic to use to rank the contributors.")
		cloudwatch_getInsightRuleReportCmd.Flags().String("period", "", "The period, in seconds, to use for the statistics in the `InsightRuleMetricDatapoint` results.")
		cloudwatch_getInsightRuleReportCmd.Flags().String("rule-name", "", "The name of the rule that you want to see data from.")
		cloudwatch_getInsightRuleReportCmd.Flags().String("start-time", "", "The start time of the data to use in the report.")
		cloudwatch_getInsightRuleReportCmd.MarkFlagRequired("end-time")
		cloudwatch_getInsightRuleReportCmd.MarkFlagRequired("period")
		cloudwatch_getInsightRuleReportCmd.MarkFlagRequired("rule-name")
		cloudwatch_getInsightRuleReportCmd.MarkFlagRequired("start-time")
	})
	cloudwatchCmd.AddCommand(cloudwatch_getInsightRuleReportCmd)
}
