package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53RecoveryReadiness_createCrossAccountAuthorizationCmd = &cobra.Command{
	Use:   "create-cross-account-authorization",
	Short: "Creates a cross-account readiness authorization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53RecoveryReadiness_createCrossAccountAuthorizationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53RecoveryReadiness_createCrossAccountAuthorizationCmd).Standalone()

		route53RecoveryReadiness_createCrossAccountAuthorizationCmd.Flags().String("cross-account-authorization", "", "The cross-account authorization.")
		route53RecoveryReadiness_createCrossAccountAuthorizationCmd.MarkFlagRequired("cross-account-authorization")
	})
	route53RecoveryReadinessCmd.AddCommand(route53RecoveryReadiness_createCrossAccountAuthorizationCmd)
}
