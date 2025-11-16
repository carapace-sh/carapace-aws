package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resiliencehub_listResourceGroupingRecommendationsCmd = &cobra.Command{
	Use:   "list-resource-grouping-recommendations",
	Short: "Lists the resource grouping recommendations suggested by Resilience Hub for your application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resiliencehub_listResourceGroupingRecommendationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(resiliencehub_listResourceGroupingRecommendationsCmd).Standalone()

		resiliencehub_listResourceGroupingRecommendationsCmd.Flags().String("app-arn", "", "Amazon Resource Name (ARN) of the Resilience Hub application.")
		resiliencehub_listResourceGroupingRecommendationsCmd.Flags().String("max-results", "", "Maximum number of grouping recommendations to be displayed per Resilience Hub application.")
		resiliencehub_listResourceGroupingRecommendationsCmd.Flags().String("next-token", "", "Null, or the token from a previous call to get the next set of results.")
	})
	resiliencehubCmd.AddCommand(resiliencehub_listResourceGroupingRecommendationsCmd)
}
