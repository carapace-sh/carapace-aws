package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var computeOptimizer_exportEcsserviceRecommendationsCmd = &cobra.Command{
	Use:   "export-ecsservice-recommendations",
	Short: "Exports optimization recommendations for Amazon ECS services on Fargate.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(computeOptimizer_exportEcsserviceRecommendationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(computeOptimizer_exportEcsserviceRecommendationsCmd).Standalone()

		computeOptimizer_exportEcsserviceRecommendationsCmd.Flags().String("account-ids", "", "The Amazon Web Services account IDs for the export Amazon ECS service recommendations.")
		computeOptimizer_exportEcsserviceRecommendationsCmd.Flags().String("fields-to-export", "", "The recommendations data to include in the export file.")
		computeOptimizer_exportEcsserviceRecommendationsCmd.Flags().String("file-format", "", "The format of the export file.")
		computeOptimizer_exportEcsserviceRecommendationsCmd.Flags().String("filters", "", "An array of objects to specify a filter that exports a more specific set of Amazon ECS service recommendations.")
		computeOptimizer_exportEcsserviceRecommendationsCmd.Flags().String("include-member-accounts", "", "If your account is the management account or the delegated administrator of an organization, this parameter indicates whether to include recommendations for resources in all member accounts of the organization.")
		computeOptimizer_exportEcsserviceRecommendationsCmd.Flags().String("s3-destination-config", "", "")
		computeOptimizer_exportEcsserviceRecommendationsCmd.MarkFlagRequired("s3-destination-config")
	})
	computeOptimizerCmd.AddCommand(computeOptimizer_exportEcsserviceRecommendationsCmd)
}
