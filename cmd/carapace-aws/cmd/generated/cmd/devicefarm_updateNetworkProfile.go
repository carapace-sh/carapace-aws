package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_updateNetworkProfileCmd = &cobra.Command{
	Use:   "update-network-profile",
	Short: "Updates the network profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_updateNetworkProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(devicefarm_updateNetworkProfileCmd).Standalone()

		devicefarm_updateNetworkProfileCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the project for which you want to update network profile settings.")
		devicefarm_updateNetworkProfileCmd.Flags().String("description", "", "The description of the network profile about which you are returning information.")
		devicefarm_updateNetworkProfileCmd.Flags().String("downlink-bandwidth-bits", "", "The data throughput rate in bits per second, as an integer from 0 to 104857600.")
		devicefarm_updateNetworkProfileCmd.Flags().String("downlink-delay-ms", "", "Delay time for all packets to destination in milliseconds as an integer from 0 to 2000.")
		devicefarm_updateNetworkProfileCmd.Flags().String("downlink-jitter-ms", "", "Time variation in the delay of received packets in milliseconds as an integer from 0 to 2000.")
		devicefarm_updateNetworkProfileCmd.Flags().String("downlink-loss-percent", "", "Proportion of received packets that fail to arrive from 0 to 100 percent.")
		devicefarm_updateNetworkProfileCmd.Flags().String("name", "", "The name of the network profile about which you are returning information.")
		devicefarm_updateNetworkProfileCmd.Flags().String("type", "", "The type of network profile to return information about.")
		devicefarm_updateNetworkProfileCmd.Flags().String("uplink-bandwidth-bits", "", "The data throughput rate in bits per second, as an integer from 0 to 104857600.")
		devicefarm_updateNetworkProfileCmd.Flags().String("uplink-delay-ms", "", "Delay time for all packets to destination in milliseconds as an integer from 0 to 2000.")
		devicefarm_updateNetworkProfileCmd.Flags().String("uplink-jitter-ms", "", "Time variation in the delay of received packets in milliseconds as an integer from 0 to 2000.")
		devicefarm_updateNetworkProfileCmd.Flags().String("uplink-loss-percent", "", "Proportion of transmitted packets that fail to arrive from 0 to 100 percent.")
		devicefarm_updateNetworkProfileCmd.MarkFlagRequired("arn")
	})
	devicefarmCmd.AddCommand(devicefarm_updateNetworkProfileCmd)
}
