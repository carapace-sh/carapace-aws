package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53RecoveryReadiness_deleteReadinessCheckCmd = &cobra.Command{
	Use:   "delete-readiness-check",
	Short: "Deletes a readiness check.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53RecoveryReadiness_deleteReadinessCheckCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53RecoveryReadiness_deleteReadinessCheckCmd).Standalone()

		route53RecoveryReadiness_deleteReadinessCheckCmd.Flags().String("readiness-check-name", "", "Name of a readiness check.")
		route53RecoveryReadiness_deleteReadinessCheckCmd.MarkFlagRequired("readiness-check-name")
	})
	route53RecoveryReadinessCmd.AddCommand(route53RecoveryReadiness_deleteReadinessCheckCmd)
}
