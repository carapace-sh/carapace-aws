package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var computeOptimizer_exportRdsdatabaseRecommendationsCmd = &cobra.Command{
	Use:   "export-rdsdatabase-recommendations",
	Short: "Export optimization recommendations for your Amazon Aurora and Amazon Relational Database Service (Amazon RDS) databases.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(computeOptimizer_exportRdsdatabaseRecommendationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(computeOptimizer_exportRdsdatabaseRecommendationsCmd).Standalone()

		computeOptimizer_exportRdsdatabaseRecommendationsCmd.Flags().String("account-ids", "", "The Amazon Web Services account IDs for the export Amazon Aurora and RDS database recommendations.")
		computeOptimizer_exportRdsdatabaseRecommendationsCmd.Flags().String("fields-to-export", "", "The recommendations data to include in the export file.")
		computeOptimizer_exportRdsdatabaseRecommendationsCmd.Flags().String("file-format", "", "The format of the export file.")
		computeOptimizer_exportRdsdatabaseRecommendationsCmd.Flags().String("filters", "", "An array of objects to specify a filter that exports a more specific set of Amazon Aurora and RDS recommendations.")
		computeOptimizer_exportRdsdatabaseRecommendationsCmd.Flags().String("include-member-accounts", "", "If your account is the management account or the delegated administrator of an organization, this parameter indicates whether to include recommendations for resources in all member accounts of the organization.")
		computeOptimizer_exportRdsdatabaseRecommendationsCmd.Flags().String("recommendation-preferences", "", "")
		computeOptimizer_exportRdsdatabaseRecommendationsCmd.Flags().String("s3-destination-config", "", "")
		computeOptimizer_exportRdsdatabaseRecommendationsCmd.MarkFlagRequired("s3-destination-config")
	})
	computeOptimizerCmd.AddCommand(computeOptimizer_exportRdsdatabaseRecommendationsCmd)
}
