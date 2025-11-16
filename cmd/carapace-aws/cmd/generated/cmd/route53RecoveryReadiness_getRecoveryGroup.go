package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53RecoveryReadiness_getRecoveryGroupCmd = &cobra.Command{
	Use:   "get-recovery-group",
	Short: "Gets details about a recovery group, including a list of the cells that are included in it.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53RecoveryReadiness_getRecoveryGroupCmd).Standalone()

	route53RecoveryReadiness_getRecoveryGroupCmd.Flags().String("recovery-group-name", "", "The name of a recovery group.")
	route53RecoveryReadiness_getRecoveryGroupCmd.MarkFlagRequired("recovery-group-name")
	route53RecoveryReadinessCmd.AddCommand(route53RecoveryReadiness_getRecoveryGroupCmd)
}
