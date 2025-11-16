package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmonitorCmd = &cobra.Command{
	Use:   "networkmonitor",
	Short: "Amazon CloudWatch Network Monitor is an Amazon Web Services active network monitoring service that identifies if a network issues exists within the Amazon Web Services network or your own company network.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmonitorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkmonitorCmd).Standalone()

	})
	rootCmd.AddCommand(networkmonitorCmd)
}
