package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_updateMetricConfigurationCmd = &cobra.Command{
	Use:   "update-metric-configuration",
	Short: "Update the summary metric configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_updateMetricConfigurationCmd).Standalone()

	iotwireless_updateMetricConfigurationCmd.Flags().String("summary-metric", "", "The value to be used to set summary metric configuration.")
	iotwirelessCmd.AddCommand(iotwireless_updateMetricConfigurationCmd)
}
