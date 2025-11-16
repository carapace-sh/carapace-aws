package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53RecoveryReadiness_createRecoveryGroupCmd = &cobra.Command{
	Use:   "create-recovery-group",
	Short: "Creates a recovery group in an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53RecoveryReadiness_createRecoveryGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53RecoveryReadiness_createRecoveryGroupCmd).Standalone()

		route53RecoveryReadiness_createRecoveryGroupCmd.Flags().String("cells", "", "A list of the cell Amazon Resource Names (ARNs) in the recovery group.")
		route53RecoveryReadiness_createRecoveryGroupCmd.Flags().String("recovery-group-name", "", "The name of the recovery group to create.")
		route53RecoveryReadiness_createRecoveryGroupCmd.Flags().String("tags", "", "")
		route53RecoveryReadiness_createRecoveryGroupCmd.MarkFlagRequired("recovery-group-name")
	})
	route53RecoveryReadinessCmd.AddCommand(route53RecoveryReadiness_createRecoveryGroupCmd)
}
