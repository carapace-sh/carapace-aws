package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53RecoveryReadiness_deleteCrossAccountAuthorizationCmd = &cobra.Command{
	Use:   "delete-cross-account-authorization",
	Short: "Deletes cross account readiness authorization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53RecoveryReadiness_deleteCrossAccountAuthorizationCmd).Standalone()

	route53RecoveryReadiness_deleteCrossAccountAuthorizationCmd.Flags().String("cross-account-authorization", "", "The cross-account authorization.")
	route53RecoveryReadiness_deleteCrossAccountAuthorizationCmd.MarkFlagRequired("cross-account-authorization")
	route53RecoveryReadinessCmd.AddCommand(route53RecoveryReadiness_deleteCrossAccountAuthorizationCmd)
}
