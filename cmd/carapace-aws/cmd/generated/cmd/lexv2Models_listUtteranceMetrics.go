package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_listUtteranceMetricsCmd = &cobra.Command{
	Use:   "list-utterance-metrics",
	Short: "To use this API operation, your IAM role must have permissions to perform the [ListAggregatedUtterances](https://docs.aws.amazon.com/lexv2/latest/APIReference/API_ListAggregatedUtterances.html) operation, which provides access to utterance-related analytics.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_listUtteranceMetricsCmd).Standalone()

	lexv2Models_listUtteranceMetricsCmd.Flags().String("attributes", "", "A list containing attributes related to the utterance that you want the response to return.")
	lexv2Models_listUtteranceMetricsCmd.Flags().String("bin-by", "", "A list of objects, each of which contains specifications for organizing the results by time.")
	lexv2Models_listUtteranceMetricsCmd.Flags().String("bot-id", "", "The identifier for the bot for which you want to retrieve utterance metrics.")
	lexv2Models_listUtteranceMetricsCmd.Flags().String("end-date-time", "", "The date and time that marks the end of the range of time for which you want to see utterance metrics.")
	lexv2Models_listUtteranceMetricsCmd.Flags().String("filters", "", "A list of objects, each of which describes a condition by which you want to filter the results.")
	lexv2Models_listUtteranceMetricsCmd.Flags().String("group-by", "", "A list of objects, each of which specifies how to group the results.")
	lexv2Models_listUtteranceMetricsCmd.Flags().String("max-results", "", "The maximum number of results to return in each page of results.")
	lexv2Models_listUtteranceMetricsCmd.Flags().String("metrics", "", "A list of objects, each of which contains a metric you want to list, the statistic for the metric you want to return, and the method by which to organize the results.")
	lexv2Models_listUtteranceMetricsCmd.Flags().String("next-token", "", "If the response from the ListUtteranceMetrics operation contains more results than specified in the maxResults parameter, a token is returned in the response.")
	lexv2Models_listUtteranceMetricsCmd.Flags().String("start-date-time", "", "The date and time that marks the beginning of the range of time for which you want to see utterance metrics.")
	lexv2Models_listUtteranceMetricsCmd.MarkFlagRequired("bot-id")
	lexv2Models_listUtteranceMetricsCmd.MarkFlagRequired("end-date-time")
	lexv2Models_listUtteranceMetricsCmd.MarkFlagRequired("metrics")
	lexv2Models_listUtteranceMetricsCmd.MarkFlagRequired("start-date-time")
	lexv2ModelsCmd.AddCommand(lexv2Models_listUtteranceMetricsCmd)
}
