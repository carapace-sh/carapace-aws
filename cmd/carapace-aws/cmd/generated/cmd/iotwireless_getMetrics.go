package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_getMetricsCmd = &cobra.Command{
	Use:   "get-metrics",
	Short: "Get the summary metrics for this AWS account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_getMetricsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotwireless_getMetricsCmd).Standalone()

		iotwireless_getMetricsCmd.Flags().String("summary-metric-queries", "", "The list of queries to retrieve the summary metrics.")
	})
	iotwirelessCmd.AddCommand(iotwireless_getMetricsCmd)
}
