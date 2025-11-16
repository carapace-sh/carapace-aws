package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var computeOptimizer_exportIdleRecommendationsCmd = &cobra.Command{
	Use:   "export-idle-recommendations",
	Short: "Export optimization recommendations for your idle resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(computeOptimizer_exportIdleRecommendationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(computeOptimizer_exportIdleRecommendationsCmd).Standalone()

		computeOptimizer_exportIdleRecommendationsCmd.Flags().String("account-ids", "", "The Amazon Web Services account IDs for the export idle resource recommendations.")
		computeOptimizer_exportIdleRecommendationsCmd.Flags().String("fields-to-export", "", "The recommendations data to include in the export file.")
		computeOptimizer_exportIdleRecommendationsCmd.Flags().String("file-format", "", "The format of the export file.")
		computeOptimizer_exportIdleRecommendationsCmd.Flags().String("filters", "", "An array of objects to specify a filter that exports a more specific set of idle resource recommendations.")
		computeOptimizer_exportIdleRecommendationsCmd.Flags().String("include-member-accounts", "", "If your account is the management account or the delegated administrator of an organization, this parameter indicates whether to include recommendations for resources in all member accounts of the organization.")
		computeOptimizer_exportIdleRecommendationsCmd.Flags().String("s3-destination-config", "", "")
		computeOptimizer_exportIdleRecommendationsCmd.MarkFlagRequired("s3-destination-config")
	})
	computeOptimizerCmd.AddCommand(computeOptimizer_exportIdleRecommendationsCmd)
}
