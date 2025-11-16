package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resiliencehub_describeResourceGroupingRecommendationTaskCmd = &cobra.Command{
	Use:   "describe-resource-grouping-recommendation-task",
	Short: "Describes the resource grouping recommendation tasks run by Resilience Hub for your application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resiliencehub_describeResourceGroupingRecommendationTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(resiliencehub_describeResourceGroupingRecommendationTaskCmd).Standalone()

		resiliencehub_describeResourceGroupingRecommendationTaskCmd.Flags().String("app-arn", "", "Amazon Resource Name (ARN) of the Resilience Hub application.")
		resiliencehub_describeResourceGroupingRecommendationTaskCmd.Flags().String("grouping-id", "", "Identifier of the grouping recommendation task.")
		resiliencehub_describeResourceGroupingRecommendationTaskCmd.MarkFlagRequired("app-arn")
	})
	resiliencehubCmd.AddCommand(resiliencehub_describeResourceGroupingRecommendationTaskCmd)
}
