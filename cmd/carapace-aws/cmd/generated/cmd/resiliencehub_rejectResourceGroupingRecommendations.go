package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resiliencehub_rejectResourceGroupingRecommendationsCmd = &cobra.Command{
	Use:   "reject-resource-grouping-recommendations",
	Short: "Rejects resource grouping recommendations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resiliencehub_rejectResourceGroupingRecommendationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(resiliencehub_rejectResourceGroupingRecommendationsCmd).Standalone()

		resiliencehub_rejectResourceGroupingRecommendationsCmd.Flags().String("app-arn", "", "Amazon Resource Name (ARN) of the Resilience Hub application.")
		resiliencehub_rejectResourceGroupingRecommendationsCmd.Flags().String("entries", "", "List of resource grouping recommendations you have selected to exclude from your application.")
		resiliencehub_rejectResourceGroupingRecommendationsCmd.MarkFlagRequired("app-arn")
		resiliencehub_rejectResourceGroupingRecommendationsCmd.MarkFlagRequired("entries")
	})
	resiliencehubCmd.AddCommand(resiliencehub_rejectResourceGroupingRecommendationsCmd)
}
