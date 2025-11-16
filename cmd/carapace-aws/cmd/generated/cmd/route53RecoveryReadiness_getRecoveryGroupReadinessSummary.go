package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53RecoveryReadiness_getRecoveryGroupReadinessSummaryCmd = &cobra.Command{
	Use:   "get-recovery-group-readiness-summary",
	Short: "Displays a summary of information about a recovery group's readiness status.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53RecoveryReadiness_getRecoveryGroupReadinessSummaryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53RecoveryReadiness_getRecoveryGroupReadinessSummaryCmd).Standalone()

		route53RecoveryReadiness_getRecoveryGroupReadinessSummaryCmd.Flags().String("max-results", "", "The number of objects that you want to return with this call.")
		route53RecoveryReadiness_getRecoveryGroupReadinessSummaryCmd.Flags().String("next-token", "", "The token that identifies which batch of results you want to see.")
		route53RecoveryReadiness_getRecoveryGroupReadinessSummaryCmd.Flags().String("recovery-group-name", "", "The name of a recovery group.")
		route53RecoveryReadiness_getRecoveryGroupReadinessSummaryCmd.MarkFlagRequired("recovery-group-name")
	})
	route53RecoveryReadinessCmd.AddCommand(route53RecoveryReadiness_getRecoveryGroupReadinessSummaryCmd)
}
