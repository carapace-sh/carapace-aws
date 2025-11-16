package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53RecoveryReadiness_getArchitectureRecommendationsCmd = &cobra.Command{
	Use:   "get-architecture-recommendations",
	Short: "Gets recommendations about architecture designs for improving resiliency for an application, based on a recovery group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53RecoveryReadiness_getArchitectureRecommendationsCmd).Standalone()

	route53RecoveryReadiness_getArchitectureRecommendationsCmd.Flags().String("max-results", "", "The number of objects that you want to return with this call.")
	route53RecoveryReadiness_getArchitectureRecommendationsCmd.Flags().String("next-token", "", "The token that identifies which batch of results you want to see.")
	route53RecoveryReadiness_getArchitectureRecommendationsCmd.Flags().String("recovery-group-name", "", "The name of a recovery group.")
	route53RecoveryReadiness_getArchitectureRecommendationsCmd.MarkFlagRequired("recovery-group-name")
	route53RecoveryReadinessCmd.AddCommand(route53RecoveryReadiness_getArchitectureRecommendationsCmd)
}
