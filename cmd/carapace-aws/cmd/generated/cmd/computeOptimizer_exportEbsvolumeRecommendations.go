package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var computeOptimizer_exportEbsvolumeRecommendationsCmd = &cobra.Command{
	Use:   "export-ebsvolume-recommendations",
	Short: "Exports optimization recommendations for Amazon EBS volumes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(computeOptimizer_exportEbsvolumeRecommendationsCmd).Standalone()

	computeOptimizer_exportEbsvolumeRecommendationsCmd.Flags().String("account-ids", "", "The IDs of the Amazon Web Services accounts for which to export Amazon EBS volume recommendations.")
	computeOptimizer_exportEbsvolumeRecommendationsCmd.Flags().String("fields-to-export", "", "The recommendations data to include in the export file.")
	computeOptimizer_exportEbsvolumeRecommendationsCmd.Flags().String("file-format", "", "The format of the export file.")
	computeOptimizer_exportEbsvolumeRecommendationsCmd.Flags().String("filters", "", "An array of objects to specify a filter that exports a more specific set of Amazon EBS volume recommendations.")
	computeOptimizer_exportEbsvolumeRecommendationsCmd.Flags().String("include-member-accounts", "", "Indicates whether to include recommendations for resources in all member accounts of the organization if your account is the management account of an organization.")
	computeOptimizer_exportEbsvolumeRecommendationsCmd.Flags().String("s3-destination-config", "", "")
	computeOptimizer_exportEbsvolumeRecommendationsCmd.MarkFlagRequired("s3-destination-config")
	computeOptimizerCmd.AddCommand(computeOptimizer_exportEbsvolumeRecommendationsCmd)
}
