package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resiliencehub_startResourceGroupingRecommendationTaskCmd = &cobra.Command{
	Use:   "start-resource-grouping-recommendation-task",
	Short: "Starts grouping recommendation task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resiliencehub_startResourceGroupingRecommendationTaskCmd).Standalone()

	resiliencehub_startResourceGroupingRecommendationTaskCmd.Flags().String("app-arn", "", "Amazon Resource Name (ARN) of the Resilience Hub application.")
	resiliencehub_startResourceGroupingRecommendationTaskCmd.MarkFlagRequired("app-arn")
	resiliencehubCmd.AddCommand(resiliencehub_startResourceGroupingRecommendationTaskCmd)
}
