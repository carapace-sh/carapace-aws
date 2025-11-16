package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_getWirelessGatewayFirmwareInformationCmd = &cobra.Command{
	Use:   "get-wireless-gateway-firmware-information",
	Short: "Gets the firmware version and other information about a wireless gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_getWirelessGatewayFirmwareInformationCmd).Standalone()

	iotwireless_getWirelessGatewayFirmwareInformationCmd.Flags().String("id", "", "The ID of the resource to get.")
	iotwireless_getWirelessGatewayFirmwareInformationCmd.MarkFlagRequired("id")
	iotwirelessCmd.AddCommand(iotwireless_getWirelessGatewayFirmwareInformationCmd)
}
