package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_getWirelessGatewayStatisticsCmd = &cobra.Command{
	Use:   "get-wireless-gateway-statistics",
	Short: "Gets operating information about a wireless gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_getWirelessGatewayStatisticsCmd).Standalone()

	iotwireless_getWirelessGatewayStatisticsCmd.Flags().String("wireless-gateway-id", "", "The ID of the wireless gateway for which to get the data.")
	iotwireless_getWirelessGatewayStatisticsCmd.MarkFlagRequired("wireless-gateway-id")
	iotwirelessCmd.AddCommand(iotwireless_getWirelessGatewayStatisticsCmd)
}
