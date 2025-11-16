package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var trustedadvisor_batchUpdateRecommendationResourceExclusionCmd = &cobra.Command{
	Use:   "batch-update-recommendation-resource-exclusion",
	Short: "Update one or more exclusion status for a list of recommendation resources",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(trustedadvisor_batchUpdateRecommendationResourceExclusionCmd).Standalone()

	trustedadvisor_batchUpdateRecommendationResourceExclusionCmd.Flags().String("recommendation-resource-exclusions", "", "A list of recommendation resource ARNs and exclusion status to update")
	trustedadvisor_batchUpdateRecommendationResourceExclusionCmd.MarkFlagRequired("recommendation-resource-exclusions")
	trustedadvisorCmd.AddCommand(trustedadvisor_batchUpdateRecommendationResourceExclusionCmd)
}
