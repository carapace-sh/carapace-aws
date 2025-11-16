package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_getNetworkAnalyzerConfigurationCmd = &cobra.Command{
	Use:   "get-network-analyzer-configuration",
	Short: "Get network analyzer configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_getNetworkAnalyzerConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotwireless_getNetworkAnalyzerConfigurationCmd).Standalone()

		iotwireless_getNetworkAnalyzerConfigurationCmd.Flags().String("configuration-name", "", "")
		iotwireless_getNetworkAnalyzerConfigurationCmd.MarkFlagRequired("configuration-name")
	})
	iotwirelessCmd.AddCommand(iotwireless_getNetworkAnalyzerConfigurationCmd)
}
