package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecast_listMonitorEvaluationsCmd = &cobra.Command{
	Use:   "list-monitor-evaluations",
	Short: "Returns a list of the monitoring evaluation results and predictor events collected by the monitor resource during different windows of time.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecast_listMonitorEvaluationsCmd).Standalone()

	forecast_listMonitorEvaluationsCmd.Flags().String("filters", "", "An array of filters.")
	forecast_listMonitorEvaluationsCmd.Flags().String("max-results", "", "The maximum number of monitoring results to return.")
	forecast_listMonitorEvaluationsCmd.Flags().String("monitor-arn", "", "The Amazon Resource Name (ARN) of the monitor resource to get results from.")
	forecast_listMonitorEvaluationsCmd.Flags().String("next-token", "", "If the result of the previous request was truncated, the response includes a `NextToken`.")
	forecast_listMonitorEvaluationsCmd.MarkFlagRequired("monitor-arn")
	forecastCmd.AddCommand(forecast_listMonitorEvaluationsCmd)
}
