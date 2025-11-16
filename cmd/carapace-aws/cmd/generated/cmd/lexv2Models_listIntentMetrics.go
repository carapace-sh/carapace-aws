package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_listIntentMetricsCmd = &cobra.Command{
	Use:   "list-intent-metrics",
	Short: "Retrieves summary metrics for the intents in your bot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_listIntentMetricsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexv2Models_listIntentMetricsCmd).Standalone()

		lexv2Models_listIntentMetricsCmd.Flags().String("bin-by", "", "A list of objects, each of which contains specifications for organizing the results by time.")
		lexv2Models_listIntentMetricsCmd.Flags().String("bot-id", "", "The identifier for the bot for which you want to retrieve intent metrics.")
		lexv2Models_listIntentMetricsCmd.Flags().String("end-date-time", "", "The date and time that marks the end of the range of time for which you want to see intent metrics.")
		lexv2Models_listIntentMetricsCmd.Flags().String("filters", "", "A list of objects, each of which describes a condition by which you want to filter the results.")
		lexv2Models_listIntentMetricsCmd.Flags().String("group-by", "", "A list of objects, each of which specifies how to group the results.")
		lexv2Models_listIntentMetricsCmd.Flags().String("max-results", "", "The maximum number of results to return in each page of results.")
		lexv2Models_listIntentMetricsCmd.Flags().String("metrics", "", "A list of objects, each of which contains a metric you want to list, the statistic for the metric you want to return, and the order by which to organize the results.")
		lexv2Models_listIntentMetricsCmd.Flags().String("next-token", "", "If the response from the ListIntentMetrics operation contains more results than specified in the maxResults parameter, a token is returned in the response.")
		lexv2Models_listIntentMetricsCmd.Flags().String("start-date-time", "", "The timestamp that marks the beginning of the range of time for which you want to see intent metrics.")
		lexv2Models_listIntentMetricsCmd.MarkFlagRequired("bot-id")
		lexv2Models_listIntentMetricsCmd.MarkFlagRequired("end-date-time")
		lexv2Models_listIntentMetricsCmd.MarkFlagRequired("metrics")
		lexv2Models_listIntentMetricsCmd.MarkFlagRequired("start-date-time")
	})
	lexv2ModelsCmd.AddCommand(lexv2Models_listIntentMetricsCmd)
}
