package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var computeOptimizer_exportLicenseRecommendationsCmd = &cobra.Command{
	Use:   "export-license-recommendations",
	Short: "Export optimization recommendations for your licenses.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(computeOptimizer_exportLicenseRecommendationsCmd).Standalone()

	computeOptimizer_exportLicenseRecommendationsCmd.Flags().String("account-ids", "", "The IDs of the Amazon Web Services accounts for which to export license recommendations.")
	computeOptimizer_exportLicenseRecommendationsCmd.Flags().String("fields-to-export", "", "The recommendations data to include in the export file.")
	computeOptimizer_exportLicenseRecommendationsCmd.Flags().String("file-format", "", "The format of the export file.")
	computeOptimizer_exportLicenseRecommendationsCmd.Flags().String("filters", "", "An array of objects to specify a filter that exports a more specific set of license recommendations.")
	computeOptimizer_exportLicenseRecommendationsCmd.Flags().String("include-member-accounts", "", "Indicates whether to include recommendations for resources in all member accounts of the organization if your account is the management account of an organization.")
	computeOptimizer_exportLicenseRecommendationsCmd.Flags().String("s3-destination-config", "", "")
	computeOptimizer_exportLicenseRecommendationsCmd.MarkFlagRequired("s3-destination-config")
	computeOptimizerCmd.AddCommand(computeOptimizer_exportLicenseRecommendationsCmd)
}
