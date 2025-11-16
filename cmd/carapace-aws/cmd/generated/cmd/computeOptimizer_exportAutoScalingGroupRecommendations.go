package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var computeOptimizer_exportAutoScalingGroupRecommendationsCmd = &cobra.Command{
	Use:   "export-auto-scaling-group-recommendations",
	Short: "Exports optimization recommendations for Auto Scaling groups.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(computeOptimizer_exportAutoScalingGroupRecommendationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(computeOptimizer_exportAutoScalingGroupRecommendationsCmd).Standalone()

		computeOptimizer_exportAutoScalingGroupRecommendationsCmd.Flags().String("account-ids", "", "The IDs of the Amazon Web Services accounts for which to export Auto Scaling group recommendations.")
		computeOptimizer_exportAutoScalingGroupRecommendationsCmd.Flags().String("fields-to-export", "", "The recommendations data to include in the export file.")
		computeOptimizer_exportAutoScalingGroupRecommendationsCmd.Flags().String("file-format", "", "The format of the export file.")
		computeOptimizer_exportAutoScalingGroupRecommendationsCmd.Flags().String("filters", "", "An array of objects to specify a filter that exports a more specific set of Auto Scaling group recommendations.")
		computeOptimizer_exportAutoScalingGroupRecommendationsCmd.Flags().String("include-member-accounts", "", "Indicates whether to include recommendations for resources in all member accounts of the organization if your account is the management account of an organization.")
		computeOptimizer_exportAutoScalingGroupRecommendationsCmd.Flags().String("recommendation-preferences", "", "An object to specify the preferences for the Auto Scaling group recommendations to export.")
		computeOptimizer_exportAutoScalingGroupRecommendationsCmd.Flags().String("s3-destination-config", "", "An object to specify the destination Amazon Simple Storage Service (Amazon S3) bucket name and key prefix for the export job.")
		computeOptimizer_exportAutoScalingGroupRecommendationsCmd.MarkFlagRequired("s3-destination-config")
	})
	computeOptimizerCmd.AddCommand(computeOptimizer_exportAutoScalingGroupRecommendationsCmd)
}
