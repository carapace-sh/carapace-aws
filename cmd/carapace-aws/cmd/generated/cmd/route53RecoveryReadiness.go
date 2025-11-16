package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53RecoveryReadinessCmd = &cobra.Command{
	Use:   "route53-recovery-readiness",
	Short: "Recovery readiness",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53RecoveryReadinessCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53RecoveryReadinessCmd).Standalone()

	})
	rootCmd.AddCommand(route53RecoveryReadinessCmd)
}
