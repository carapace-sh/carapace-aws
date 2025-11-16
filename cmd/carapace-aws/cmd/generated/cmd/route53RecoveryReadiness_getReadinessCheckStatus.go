package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53RecoveryReadiness_getReadinessCheckStatusCmd = &cobra.Command{
	Use:   "get-readiness-check-status",
	Short: "Gets the readiness status for an individual readiness check.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53RecoveryReadiness_getReadinessCheckStatusCmd).Standalone()

	route53RecoveryReadiness_getReadinessCheckStatusCmd.Flags().String("max-results", "", "The number of objects that you want to return with this call.")
	route53RecoveryReadiness_getReadinessCheckStatusCmd.Flags().String("next-token", "", "The token that identifies which batch of results you want to see.")
	route53RecoveryReadiness_getReadinessCheckStatusCmd.Flags().String("readiness-check-name", "", "Name of a readiness check.")
	route53RecoveryReadiness_getReadinessCheckStatusCmd.MarkFlagRequired("readiness-check-name")
	route53RecoveryReadinessCmd.AddCommand(route53RecoveryReadiness_getReadinessCheckStatusCmd)
}
