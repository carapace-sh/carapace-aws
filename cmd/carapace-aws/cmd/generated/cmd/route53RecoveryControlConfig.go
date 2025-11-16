package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53RecoveryControlConfigCmd = &cobra.Command{
	Use:   "route53-recovery-control-config",
	Short: "Recovery Control Configuration API Reference for Amazon Route 53 Application Recovery Controller",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53RecoveryControlConfigCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53RecoveryControlConfigCmd).Standalone()

	})
	rootCmd.AddCommand(route53RecoveryControlConfigCmd)
}
