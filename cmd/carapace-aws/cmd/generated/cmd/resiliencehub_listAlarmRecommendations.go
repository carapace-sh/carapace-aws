package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resiliencehub_listAlarmRecommendationsCmd = &cobra.Command{
	Use:   "list-alarm-recommendations",
	Short: "Lists the alarm recommendations for an Resilience Hub application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resiliencehub_listAlarmRecommendationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(resiliencehub_listAlarmRecommendationsCmd).Standalone()

		resiliencehub_listAlarmRecommendationsCmd.Flags().String("assessment-arn", "", "Amazon Resource Name (ARN) of the assessment.")
		resiliencehub_listAlarmRecommendationsCmd.Flags().String("max-results", "", "Maximum number of results to include in the response.")
		resiliencehub_listAlarmRecommendationsCmd.Flags().String("next-token", "", "Null, or the token from a previous call to get the next set of results.")
		resiliencehub_listAlarmRecommendationsCmd.MarkFlagRequired("assessment-arn")
	})
	resiliencehubCmd.AddCommand(resiliencehub_listAlarmRecommendationsCmd)
}
