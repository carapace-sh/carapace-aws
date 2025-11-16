package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resiliencehub_listSopRecommendationsCmd = &cobra.Command{
	Use:   "list-sop-recommendations",
	Short: "Lists the standard operating procedure (SOP) recommendations for the Resilience Hub applications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resiliencehub_listSopRecommendationsCmd).Standalone()

	resiliencehub_listSopRecommendationsCmd.Flags().String("assessment-arn", "", "Amazon Resource Name (ARN) of the assessment.")
	resiliencehub_listSopRecommendationsCmd.Flags().String("max-results", "", "Maximum number of results to include in the response.")
	resiliencehub_listSopRecommendationsCmd.Flags().String("next-token", "", "Null, or the token from a previous call to get the next set of results.")
	resiliencehub_listSopRecommendationsCmd.MarkFlagRequired("assessment-arn")
	resiliencehubCmd.AddCommand(resiliencehub_listSopRecommendationsCmd)
}
