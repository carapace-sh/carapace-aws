package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resiliencehub_listAppComponentRecommendationsCmd = &cobra.Command{
	Use:   "list-app-component-recommendations",
	Short: "Lists the recommendations for an Resilience Hub Application Component.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resiliencehub_listAppComponentRecommendationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(resiliencehub_listAppComponentRecommendationsCmd).Standalone()

		resiliencehub_listAppComponentRecommendationsCmd.Flags().String("assessment-arn", "", "Amazon Resource Name (ARN) of the assessment.")
		resiliencehub_listAppComponentRecommendationsCmd.Flags().String("max-results", "", "Maximum number of results to include in the response.")
		resiliencehub_listAppComponentRecommendationsCmd.Flags().String("next-token", "", "Null, or the token from a previous call to get the next set of results.")
		resiliencehub_listAppComponentRecommendationsCmd.MarkFlagRequired("assessment-arn")
	})
	resiliencehubCmd.AddCommand(resiliencehub_listAppComponentRecommendationsCmd)
}
