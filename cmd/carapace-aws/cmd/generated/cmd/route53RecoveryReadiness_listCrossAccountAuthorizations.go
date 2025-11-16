package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53RecoveryReadiness_listCrossAccountAuthorizationsCmd = &cobra.Command{
	Use:   "list-cross-account-authorizations",
	Short: "Lists the cross-account readiness authorizations that are in place for an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53RecoveryReadiness_listCrossAccountAuthorizationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53RecoveryReadiness_listCrossAccountAuthorizationsCmd).Standalone()

		route53RecoveryReadiness_listCrossAccountAuthorizationsCmd.Flags().String("max-results", "", "The number of objects that you want to return with this call.")
		route53RecoveryReadiness_listCrossAccountAuthorizationsCmd.Flags().String("next-token", "", "The token that identifies which batch of results you want to see.")
	})
	route53RecoveryReadinessCmd.AddCommand(route53RecoveryReadiness_listCrossAccountAuthorizationsCmd)
}
