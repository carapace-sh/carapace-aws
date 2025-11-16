package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_createNetworkAnalyzerConfigurationCmd = &cobra.Command{
	Use:   "create-network-analyzer-configuration",
	Short: "Creates a new network analyzer configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_createNetworkAnalyzerConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotwireless_createNetworkAnalyzerConfigurationCmd).Standalone()

		iotwireless_createNetworkAnalyzerConfigurationCmd.Flags().String("client-request-token", "", "")
		iotwireless_createNetworkAnalyzerConfigurationCmd.Flags().String("description", "", "")
		iotwireless_createNetworkAnalyzerConfigurationCmd.Flags().String("multicast-groups", "", "Multicast Group resources to add to the network analyzer configruation.")
		iotwireless_createNetworkAnalyzerConfigurationCmd.Flags().String("name", "", "")
		iotwireless_createNetworkAnalyzerConfigurationCmd.Flags().String("tags", "", "")
		iotwireless_createNetworkAnalyzerConfigurationCmd.Flags().String("trace-content", "", "")
		iotwireless_createNetworkAnalyzerConfigurationCmd.Flags().String("wireless-devices", "", "Wireless device resources to add to the network analyzer configuration.")
		iotwireless_createNetworkAnalyzerConfigurationCmd.Flags().String("wireless-gateways", "", "Wireless gateway resources to add to the network analyzer configuration.")
		iotwireless_createNetworkAnalyzerConfigurationCmd.MarkFlagRequired("name")
	})
	iotwirelessCmd.AddCommand(iotwireless_createNetworkAnalyzerConfigurationCmd)
}
