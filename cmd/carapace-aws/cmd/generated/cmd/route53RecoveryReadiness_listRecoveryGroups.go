package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53RecoveryReadiness_listRecoveryGroupsCmd = &cobra.Command{
	Use:   "list-recovery-groups",
	Short: "Lists the recovery groups in an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53RecoveryReadiness_listRecoveryGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53RecoveryReadiness_listRecoveryGroupsCmd).Standalone()

		route53RecoveryReadiness_listRecoveryGroupsCmd.Flags().String("max-results", "", "The number of objects that you want to return with this call.")
		route53RecoveryReadiness_listRecoveryGroupsCmd.Flags().String("next-token", "", "The token that identifies which batch of results you want to see.")
	})
	route53RecoveryReadinessCmd.AddCommand(route53RecoveryReadiness_listRecoveryGroupsCmd)
}
