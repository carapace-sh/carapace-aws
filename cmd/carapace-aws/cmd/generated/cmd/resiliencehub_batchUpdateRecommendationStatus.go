package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resiliencehub_batchUpdateRecommendationStatusCmd = &cobra.Command{
	Use:   "batch-update-recommendation-status",
	Short: "Enables you to include or exclude one or more operational recommendations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resiliencehub_batchUpdateRecommendationStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(resiliencehub_batchUpdateRecommendationStatusCmd).Standalone()

		resiliencehub_batchUpdateRecommendationStatusCmd.Flags().String("app-arn", "", "Amazon Resource Name (ARN) of the Resilience Hub application.")
		resiliencehub_batchUpdateRecommendationStatusCmd.Flags().String("request-entries", "", "Defines the list of operational recommendations that need to be included or excluded.")
		resiliencehub_batchUpdateRecommendationStatusCmd.MarkFlagRequired("app-arn")
		resiliencehub_batchUpdateRecommendationStatusCmd.MarkFlagRequired("request-entries")
	})
	resiliencehubCmd.AddCommand(resiliencehub_batchUpdateRecommendationStatusCmd)
}
