package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resiliencehub_acceptResourceGroupingRecommendationsCmd = &cobra.Command{
	Use:   "accept-resource-grouping-recommendations",
	Short: "Accepts the resource grouping recommendations suggested by Resilience Hub for your application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resiliencehub_acceptResourceGroupingRecommendationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(resiliencehub_acceptResourceGroupingRecommendationsCmd).Standalone()

		resiliencehub_acceptResourceGroupingRecommendationsCmd.Flags().String("app-arn", "", "Amazon Resource Name (ARN) of the Resilience Hub application.")
		resiliencehub_acceptResourceGroupingRecommendationsCmd.Flags().String("entries", "", "List of resource grouping recommendations you want to include in your application.")
		resiliencehub_acceptResourceGroupingRecommendationsCmd.MarkFlagRequired("app-arn")
		resiliencehub_acceptResourceGroupingRecommendationsCmd.MarkFlagRequired("entries")
	})
	resiliencehubCmd.AddCommand(resiliencehub_acceptResourceGroupingRecommendationsCmd)
}
