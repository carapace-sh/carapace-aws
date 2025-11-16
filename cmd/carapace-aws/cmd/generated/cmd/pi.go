package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var piCmd = &cobra.Command{
	Use:   "pi",
	Short: "Amazon RDS Performance Insights\n\nAmazon RDS Performance Insights enables you to monitor and explore different dimensions of database load based on data captured from a running DB instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(piCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(piCmd).Standalone()

	})
	rootCmd.AddCommand(piCmd)
}
