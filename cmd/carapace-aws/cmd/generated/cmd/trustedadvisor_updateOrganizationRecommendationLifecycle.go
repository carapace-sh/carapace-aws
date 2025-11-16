package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var trustedadvisor_updateOrganizationRecommendationLifecycleCmd = &cobra.Command{
	Use:   "update-organization-recommendation-lifecycle",
	Short: "Update the lifecycle of a Recommendation within an Organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(trustedadvisor_updateOrganizationRecommendationLifecycleCmd).Standalone()

	trustedadvisor_updateOrganizationRecommendationLifecycleCmd.Flags().String("lifecycle-stage", "", "The new lifecycle stage")
	trustedadvisor_updateOrganizationRecommendationLifecycleCmd.Flags().String("organization-recommendation-identifier", "", "The Recommendation identifier for AWS Trusted Advisor Priority recommendations")
	trustedadvisor_updateOrganizationRecommendationLifecycleCmd.Flags().String("update-reason", "", "Reason for the lifecycle stage change")
	trustedadvisor_updateOrganizationRecommendationLifecycleCmd.Flags().String("update-reason-code", "", "Reason code for the lifecycle state change")
	trustedadvisor_updateOrganizationRecommendationLifecycleCmd.MarkFlagRequired("lifecycle-stage")
	trustedadvisor_updateOrganizationRecommendationLifecycleCmd.MarkFlagRequired("organization-recommendation-identifier")
	trustedadvisorCmd.AddCommand(trustedadvisor_updateOrganizationRecommendationLifecycleCmd)
}
