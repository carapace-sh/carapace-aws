package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_createNetworkProfileCmd = &cobra.Command{
	Use:   "create-network-profile",
	Short: "Creates a network profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_createNetworkProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(devicefarm_createNetworkProfileCmd).Standalone()

		devicefarm_createNetworkProfileCmd.Flags().String("description", "", "The description of the network profile.")
		devicefarm_createNetworkProfileCmd.Flags().String("downlink-bandwidth-bits", "", "The data throughput rate in bits per second, as an integer from 0 to 104857600.")
		devicefarm_createNetworkProfileCmd.Flags().String("downlink-delay-ms", "", "Delay time for all packets to destination in milliseconds as an integer from 0 to 2000.")
		devicefarm_createNetworkProfileCmd.Flags().String("downlink-jitter-ms", "", "Time variation in the delay of received packets in milliseconds as an integer from 0 to 2000.")
		devicefarm_createNetworkProfileCmd.Flags().String("downlink-loss-percent", "", "Proportion of received packets that fail to arrive from 0 to 100 percent.")
		devicefarm_createNetworkProfileCmd.Flags().String("name", "", "The name for the new network profile.")
		devicefarm_createNetworkProfileCmd.Flags().String("project-arn", "", "The Amazon Resource Name (ARN) of the project for which you want to create a network profile.")
		devicefarm_createNetworkProfileCmd.Flags().String("type", "", "The type of network profile to create.")
		devicefarm_createNetworkProfileCmd.Flags().String("uplink-bandwidth-bits", "", "The data throughput rate in bits per second, as an integer from 0 to 104857600.")
		devicefarm_createNetworkProfileCmd.Flags().String("uplink-delay-ms", "", "Delay time for all packets to destination in milliseconds as an integer from 0 to 2000.")
		devicefarm_createNetworkProfileCmd.Flags().String("uplink-jitter-ms", "", "Time variation in the delay of received packets in milliseconds as an integer from 0 to 2000.")
		devicefarm_createNetworkProfileCmd.Flags().String("uplink-loss-percent", "", "Proportion of transmitted packets that fail to arrive from 0 to 100 percent.")
		devicefarm_createNetworkProfileCmd.MarkFlagRequired("name")
		devicefarm_createNetworkProfileCmd.MarkFlagRequired("project-arn")
	})
	devicefarmCmd.AddCommand(devicefarm_createNetworkProfileCmd)
}
