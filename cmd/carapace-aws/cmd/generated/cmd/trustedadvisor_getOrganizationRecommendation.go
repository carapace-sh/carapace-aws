package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var trustedadvisor_getOrganizationRecommendationCmd = &cobra.Command{
	Use:   "get-organization-recommendation",
	Short: "Get a specific recommendation within an AWS Organizations organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(trustedadvisor_getOrganizationRecommendationCmd).Standalone()

	trustedadvisor_getOrganizationRecommendationCmd.Flags().String("organization-recommendation-identifier", "", "The Recommendation identifier")
	trustedadvisor_getOrganizationRecommendationCmd.MarkFlagRequired("organization-recommendation-identifier")
	trustedadvisorCmd.AddCommand(trustedadvisor_getOrganizationRecommendationCmd)
}
