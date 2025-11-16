package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53RecoveryReadiness_updateRecoveryGroupCmd = &cobra.Command{
	Use:   "update-recovery-group",
	Short: "Updates a recovery group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53RecoveryReadiness_updateRecoveryGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53RecoveryReadiness_updateRecoveryGroupCmd).Standalone()

		route53RecoveryReadiness_updateRecoveryGroupCmd.Flags().String("cells", "", "A list of cell Amazon Resource Names (ARNs).")
		route53RecoveryReadiness_updateRecoveryGroupCmd.Flags().String("recovery-group-name", "", "The name of a recovery group.")
		route53RecoveryReadiness_updateRecoveryGroupCmd.MarkFlagRequired("cells")
		route53RecoveryReadiness_updateRecoveryGroupCmd.MarkFlagRequired("recovery-group-name")
	})
	route53RecoveryReadinessCmd.AddCommand(route53RecoveryReadiness_updateRecoveryGroupCmd)
}
