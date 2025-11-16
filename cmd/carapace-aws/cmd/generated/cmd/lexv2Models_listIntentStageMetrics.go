package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_listIntentStageMetricsCmd = &cobra.Command{
	Use:   "list-intent-stage-metrics",
	Short: "Retrieves summary metrics for the stages within intents in your bot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_listIntentStageMetricsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexv2Models_listIntentStageMetricsCmd).Standalone()

		lexv2Models_listIntentStageMetricsCmd.Flags().String("bin-by", "", "A list of objects, each of which contains specifications for organizing the results by time.")
		lexv2Models_listIntentStageMetricsCmd.Flags().String("bot-id", "", "The identifier for the bot for which you want to retrieve intent stage metrics.")
		lexv2Models_listIntentStageMetricsCmd.Flags().String("end-date-time", "", "The date and time that marks the end of the range of time for which you want to see intent stage metrics.")
		lexv2Models_listIntentStageMetricsCmd.Flags().String("filters", "", "A list of objects, each of which describes a condition by which you want to filter the results.")
		lexv2Models_listIntentStageMetricsCmd.Flags().String("group-by", "", "A list of objects, each of which specifies how to group the results.")
		lexv2Models_listIntentStageMetricsCmd.Flags().String("max-results", "", "The maximum number of results to return in each page of results.")
		lexv2Models_listIntentStageMetricsCmd.Flags().String("metrics", "", "A list of objects, each of which contains a metric you want to list, the statistic for the metric you want to return, and the method by which to organize the results.")
		lexv2Models_listIntentStageMetricsCmd.Flags().String("next-token", "", "If the response from the ListIntentStageMetrics operation contains more results than specified in the maxResults parameter, a token is returned in the response.")
		lexv2Models_listIntentStageMetricsCmd.Flags().String("start-date-time", "", "The date and time that marks the beginning of the range of time for which you want to see intent stage metrics.")
		lexv2Models_listIntentStageMetricsCmd.MarkFlagRequired("bot-id")
		lexv2Models_listIntentStageMetricsCmd.MarkFlagRequired("end-date-time")
		lexv2Models_listIntentStageMetricsCmd.MarkFlagRequired("metrics")
		lexv2Models_listIntentStageMetricsCmd.MarkFlagRequired("start-date-time")
	})
	lexv2ModelsCmd.AddCommand(lexv2Models_listIntentStageMetricsCmd)
}
