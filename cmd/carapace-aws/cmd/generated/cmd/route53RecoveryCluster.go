package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53RecoveryClusterCmd = &cobra.Command{
	Use:   "route53-recovery-cluster",
	Short: "Welcome to the Routing Control (Recovery Cluster) API Reference Guide for Amazon Route 53 Application Recovery Controller.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53RecoveryClusterCmd).Standalone()

	rootCmd.AddCommand(route53RecoveryClusterCmd)
}
