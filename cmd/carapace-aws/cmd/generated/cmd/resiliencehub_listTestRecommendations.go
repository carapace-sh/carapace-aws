package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resiliencehub_listTestRecommendationsCmd = &cobra.Command{
	Use:   "list-test-recommendations",
	Short: "Lists the test recommendations for the Resilience Hub application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resiliencehub_listTestRecommendationsCmd).Standalone()

	resiliencehub_listTestRecommendationsCmd.Flags().String("assessment-arn", "", "Amazon Resource Name (ARN) of the assessment.")
	resiliencehub_listTestRecommendationsCmd.Flags().String("max-results", "", "Maximum number of results to include in the response.")
	resiliencehub_listTestRecommendationsCmd.Flags().String("next-token", "", "Null, or the token from a previous call to get the next set of results.")
	resiliencehub_listTestRecommendationsCmd.MarkFlagRequired("assessment-arn")
	resiliencehubCmd.AddCommand(resiliencehub_listTestRecommendationsCmd)
}
