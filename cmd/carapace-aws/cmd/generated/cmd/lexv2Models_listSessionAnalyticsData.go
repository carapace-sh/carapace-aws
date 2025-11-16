package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_listSessionAnalyticsDataCmd = &cobra.Command{
	Use:   "list-session-analytics-data",
	Short: "Retrieves a list of metadata for individual user sessions with your bot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_listSessionAnalyticsDataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexv2Models_listSessionAnalyticsDataCmd).Standalone()

		lexv2Models_listSessionAnalyticsDataCmd.Flags().String("bot-id", "", "The identifier for the bot for which you want to retrieve session analytics.")
		lexv2Models_listSessionAnalyticsDataCmd.Flags().String("end-date-time", "", "The date and time that marks the end of the range of time for which you want to see session analytics.")
		lexv2Models_listSessionAnalyticsDataCmd.Flags().String("filters", "", "A list of objects, each of which describes a condition by which you want to filter the results.")
		lexv2Models_listSessionAnalyticsDataCmd.Flags().String("max-results", "", "The maximum number of results to return in each page of results.")
		lexv2Models_listSessionAnalyticsDataCmd.Flags().String("next-token", "", "If the response from the ListSessionAnalyticsData operation contains more results than specified in the maxResults parameter, a token is returned in the response.")
		lexv2Models_listSessionAnalyticsDataCmd.Flags().String("sort-by", "", "An object specifying the measure and method by which to sort the session analytics data.")
		lexv2Models_listSessionAnalyticsDataCmd.Flags().String("start-date-time", "", "The date and time that marks the beginning of the range of time for which you want to see session analytics.")
		lexv2Models_listSessionAnalyticsDataCmd.MarkFlagRequired("bot-id")
		lexv2Models_listSessionAnalyticsDataCmd.MarkFlagRequired("end-date-time")
		lexv2Models_listSessionAnalyticsDataCmd.MarkFlagRequired("start-date-time")
	})
	lexv2ModelsCmd.AddCommand(lexv2Models_listSessionAnalyticsDataCmd)
}
