package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_getMetricConfigurationCmd = &cobra.Command{
	Use:   "get-metric-configuration",
	Short: "Get the metric configuration status for this AWS account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_getMetricConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotwireless_getMetricConfigurationCmd).Standalone()

	})
	iotwirelessCmd.AddCommand(iotwireless_getMetricConfigurationCmd)
}
