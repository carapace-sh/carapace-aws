package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53RecoveryReadiness_deleteRecoveryGroupCmd = &cobra.Command{
	Use:   "delete-recovery-group",
	Short: "Deletes a recovery group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53RecoveryReadiness_deleteRecoveryGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53RecoveryReadiness_deleteRecoveryGroupCmd).Standalone()

		route53RecoveryReadiness_deleteRecoveryGroupCmd.Flags().String("recovery-group-name", "", "The name of a recovery group.")
		route53RecoveryReadiness_deleteRecoveryGroupCmd.MarkFlagRequired("recovery-group-name")
	})
	route53RecoveryReadinessCmd.AddCommand(route53RecoveryReadiness_deleteRecoveryGroupCmd)
}
