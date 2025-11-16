package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var internetmonitorCmd = &cobra.Command{
	Use:   "internetmonitor",
	Short: "Amazon CloudWatch Internet Monitor provides visibility into how internet issues impact the performance and availability between your applications hosted on Amazon Web Services and your end users.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(internetmonitorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(internetmonitorCmd).Standalone()

	})
	rootCmd.AddCommand(internetmonitorCmd)
}
