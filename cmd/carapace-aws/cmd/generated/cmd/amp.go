package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ampCmd = &cobra.Command{
	Use:   "amp",
	Short: "Amazon Managed Service for Prometheus is a serverless, Prometheus-compatible monitoring service for container metrics that makes it easier to securely monitor container environments at scale.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ampCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ampCmd).Standalone()

	})
	rootCmd.AddCommand(ampCmd)
}
