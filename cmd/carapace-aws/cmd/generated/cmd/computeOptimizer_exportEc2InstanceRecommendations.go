package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var computeOptimizer_exportEc2InstanceRecommendationsCmd = &cobra.Command{
	Use:   "export-ec2-instance-recommendations",
	Short: "Exports optimization recommendations for Amazon EC2 instances.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(computeOptimizer_exportEc2InstanceRecommendationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(computeOptimizer_exportEc2InstanceRecommendationsCmd).Standalone()

		computeOptimizer_exportEc2InstanceRecommendationsCmd.Flags().String("account-ids", "", "The IDs of the Amazon Web Services accounts for which to export instance recommendations.")
		computeOptimizer_exportEc2InstanceRecommendationsCmd.Flags().String("fields-to-export", "", "The recommendations data to include in the export file.")
		computeOptimizer_exportEc2InstanceRecommendationsCmd.Flags().String("file-format", "", "The format of the export file.")
		computeOptimizer_exportEc2InstanceRecommendationsCmd.Flags().String("filters", "", "An array of objects to specify a filter that exports a more specific set of instance recommendations.")
		computeOptimizer_exportEc2InstanceRecommendationsCmd.Flags().String("include-member-accounts", "", "Indicates whether to include recommendations for resources in all member accounts of the organization if your account is the management account of an organization.")
		computeOptimizer_exportEc2InstanceRecommendationsCmd.Flags().String("recommendation-preferences", "", "An object to specify the preferences for the Amazon EC2 instance recommendations to export.")
		computeOptimizer_exportEc2InstanceRecommendationsCmd.Flags().String("s3-destination-config", "", "An object to specify the destination Amazon Simple Storage Service (Amazon S3) bucket name and key prefix for the export job.")
		computeOptimizer_exportEc2InstanceRecommendationsCmd.MarkFlagRequired("s3-destination-config")
	})
	computeOptimizerCmd.AddCommand(computeOptimizer_exportEc2InstanceRecommendationsCmd)
}
