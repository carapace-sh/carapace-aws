package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudwatch_listMetricsCmd = &cobra.Command{
	Use:   "list-metrics",
	Short: "List the specified metrics.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudwatch_listMetricsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudwatch_listMetricsCmd).Standalone()

		cloudwatch_listMetricsCmd.Flags().String("dimensions", "", "The dimensions to filter against.")
		cloudwatch_listMetricsCmd.Flags().String("include-linked-accounts", "", "If you are using this operation in a monitoring account, specify `true` to include metrics from source accounts in the returned data.")
		cloudwatch_listMetricsCmd.Flags().String("metric-name", "", "The name of the metric to filter against.")
		cloudwatch_listMetricsCmd.Flags().String("namespace", "", "The metric namespace to filter against.")
		cloudwatch_listMetricsCmd.Flags().String("next-token", "", "The token returned by a previous call to indicate that there is more data available.")
		cloudwatch_listMetricsCmd.Flags().String("owning-account", "", "When you use this operation in a monitoring account, use this field to return metrics only from one source account.")
		cloudwatch_listMetricsCmd.Flags().String("recently-active", "", "To filter the results to show only metrics that have had data points published in the past three hours, specify this parameter with a value of `PT3H`.")
	})
	cloudwatchCmd.AddCommand(cloudwatch_listMetricsCmd)
}
