package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53RecoveryReadiness_listReadinessChecksCmd = &cobra.Command{
	Use:   "list-readiness-checks",
	Short: "Lists the readiness checks for an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53RecoveryReadiness_listReadinessChecksCmd).Standalone()

	route53RecoveryReadiness_listReadinessChecksCmd.Flags().String("max-results", "", "The number of objects that you want to return with this call.")
	route53RecoveryReadiness_listReadinessChecksCmd.Flags().String("next-token", "", "The token that identifies which batch of results you want to see.")
	route53RecoveryReadinessCmd.AddCommand(route53RecoveryReadiness_listReadinessChecksCmd)
}
