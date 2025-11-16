package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var trustedadvisor_getRecommendationCmd = &cobra.Command{
	Use:   "get-recommendation",
	Short: "Get a specific Recommendation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(trustedadvisor_getRecommendationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(trustedadvisor_getRecommendationCmd).Standalone()

		trustedadvisor_getRecommendationCmd.Flags().String("recommendation-identifier", "", "The Recommendation identifier")
		trustedadvisor_getRecommendationCmd.MarkFlagRequired("recommendation-identifier")
	})
	trustedadvisorCmd.AddCommand(trustedadvisor_getRecommendationCmd)
}
