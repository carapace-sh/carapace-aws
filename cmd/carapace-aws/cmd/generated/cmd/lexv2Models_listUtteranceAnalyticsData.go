package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_listUtteranceAnalyticsDataCmd = &cobra.Command{
	Use:   "list-utterance-analytics-data",
	Short: "To use this API operation, your IAM role must have permissions to perform the [ListAggregatedUtterances](https://docs.aws.amazon.com/lexv2/latest/APIReference/API_ListAggregatedUtterances.html) operation, which provides access to utterance-related analytics.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_listUtteranceAnalyticsDataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexv2Models_listUtteranceAnalyticsDataCmd).Standalone()

		lexv2Models_listUtteranceAnalyticsDataCmd.Flags().String("bot-id", "", "The identifier for the bot for which you want to retrieve utterance analytics.")
		lexv2Models_listUtteranceAnalyticsDataCmd.Flags().String("end-date-time", "", "The date and time that marks the end of the range of time for which you want to see utterance analytics.")
		lexv2Models_listUtteranceAnalyticsDataCmd.Flags().String("filters", "", "A list of objects, each of which describes a condition by which you want to filter the results.")
		lexv2Models_listUtteranceAnalyticsDataCmd.Flags().String("max-results", "", "The maximum number of results to return in each page of results.")
		lexv2Models_listUtteranceAnalyticsDataCmd.Flags().String("next-token", "", "If the response from the ListUtteranceAnalyticsData operation contains more results than specified in the maxResults parameter, a token is returned in the response.")
		lexv2Models_listUtteranceAnalyticsDataCmd.Flags().String("sort-by", "", "An object specifying the measure and method by which to sort the utterance analytics data.")
		lexv2Models_listUtteranceAnalyticsDataCmd.Flags().String("start-date-time", "", "The date and time that marks the beginning of the range of time for which you want to see utterance analytics.")
		lexv2Models_listUtteranceAnalyticsDataCmd.MarkFlagRequired("bot-id")
		lexv2Models_listUtteranceAnalyticsDataCmd.MarkFlagRequired("end-date-time")
		lexv2Models_listUtteranceAnalyticsDataCmd.MarkFlagRequired("start-date-time")
	})
	lexv2ModelsCmd.AddCommand(lexv2Models_listUtteranceAnalyticsDataCmd)
}
