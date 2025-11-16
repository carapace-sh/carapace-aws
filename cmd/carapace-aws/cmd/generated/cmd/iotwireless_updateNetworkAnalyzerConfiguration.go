package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_updateNetworkAnalyzerConfigurationCmd = &cobra.Command{
	Use:   "update-network-analyzer-configuration",
	Short: "Update network analyzer configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_updateNetworkAnalyzerConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotwireless_updateNetworkAnalyzerConfigurationCmd).Standalone()

		iotwireless_updateNetworkAnalyzerConfigurationCmd.Flags().String("configuration-name", "", "")
		iotwireless_updateNetworkAnalyzerConfigurationCmd.Flags().String("description", "", "")
		iotwireless_updateNetworkAnalyzerConfigurationCmd.Flags().String("multicast-groups-to-add", "", "Multicast group resources to add to the network analyzer configuration.")
		iotwireless_updateNetworkAnalyzerConfigurationCmd.Flags().String("multicast-groups-to-remove", "", "Multicast group resources to remove from the network analyzer configuration.")
		iotwireless_updateNetworkAnalyzerConfigurationCmd.Flags().String("trace-content", "", "")
		iotwireless_updateNetworkAnalyzerConfigurationCmd.Flags().String("wireless-devices-to-add", "", "Wireless device resources to add to the network analyzer configuration.")
		iotwireless_updateNetworkAnalyzerConfigurationCmd.Flags().String("wireless-devices-to-remove", "", "Wireless device resources to remove from the network analyzer configuration.")
		iotwireless_updateNetworkAnalyzerConfigurationCmd.Flags().String("wireless-gateways-to-add", "", "Wireless gateway resources to add to the network analyzer configuration.")
		iotwireless_updateNetworkAnalyzerConfigurationCmd.Flags().String("wireless-gateways-to-remove", "", "Wireless gateway resources to remove from the network analyzer configuration.")
		iotwireless_updateNetworkAnalyzerConfigurationCmd.MarkFlagRequired("configuration-name")
	})
	iotwirelessCmd.AddCommand(iotwireless_updateNetworkAnalyzerConfigurationCmd)
}
