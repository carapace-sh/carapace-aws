package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bcmRecommendedActions_listRecommendedActionsCmd = &cobra.Command{
	Use:   "list-recommended-actions",
	Short: "Returns a list of recommended actions that match the filter criteria.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bcmRecommendedActions_listRecommendedActionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bcmRecommendedActions_listRecommendedActionsCmd).Standalone()

		bcmRecommendedActions_listRecommendedActionsCmd.Flags().String("filter", "", "The criteria that you want all returned recommended actions to match.")
		bcmRecommendedActions_listRecommendedActionsCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
		bcmRecommendedActions_listRecommendedActionsCmd.Flags().String("next-token", "", "The pagination token that indicates the next set of results that you want to retrieve.")
	})
	bcmRecommendedActionsCmd.AddCommand(bcmRecommendedActions_listRecommendedActionsCmd)
}
