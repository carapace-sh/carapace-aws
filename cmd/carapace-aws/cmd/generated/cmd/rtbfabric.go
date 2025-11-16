package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rtbfabricCmd = &cobra.Command{
	Use:   "rtbfabric",
	Short: "Amazon Web Services RTB Fabric provides secure, low-latency infrastructure for connecting real-time bidding (RTB) applications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rtbfabricCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rtbfabricCmd).Standalone()

	})
	rootCmd.AddCommand(rtbfabricCmd)
}
