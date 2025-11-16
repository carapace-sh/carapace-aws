package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var trustedadvisor_updateRecommendationLifecycleCmd = &cobra.Command{
	Use:   "update-recommendation-lifecycle",
	Short: "Update the lifecyle of a Recommendation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(trustedadvisor_updateRecommendationLifecycleCmd).Standalone()

	trustedadvisor_updateRecommendationLifecycleCmd.Flags().String("lifecycle-stage", "", "The new lifecycle stage")
	trustedadvisor_updateRecommendationLifecycleCmd.Flags().String("recommendation-identifier", "", "The Recommendation identifier for AWS Trusted Advisor Priority recommendations")
	trustedadvisor_updateRecommendationLifecycleCmd.Flags().String("update-reason", "", "Reason for the lifecycle stage change")
	trustedadvisor_updateRecommendationLifecycleCmd.Flags().String("update-reason-code", "", "Reason code for the lifecycle state change")
	trustedadvisor_updateRecommendationLifecycleCmd.MarkFlagRequired("lifecycle-stage")
	trustedadvisor_updateRecommendationLifecycleCmd.MarkFlagRequired("recommendation-identifier")
	trustedadvisorCmd.AddCommand(trustedadvisor_updateRecommendationLifecycleCmd)
}
