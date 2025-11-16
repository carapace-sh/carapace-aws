package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_disassociateWirelessDeviceFromFuotaTaskCmd = &cobra.Command{
	Use:   "disassociate-wireless-device-from-fuota-task",
	Short: "Disassociates a wireless device from a FUOTA task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_disassociateWirelessDeviceFromFuotaTaskCmd).Standalone()

	iotwireless_disassociateWirelessDeviceFromFuotaTaskCmd.Flags().String("id", "", "")
	iotwireless_disassociateWirelessDeviceFromFuotaTaskCmd.Flags().String("wireless-device-id", "", "")
	iotwireless_disassociateWirelessDeviceFromFuotaTaskCmd.MarkFlagRequired("id")
	iotwireless_disassociateWirelessDeviceFromFuotaTaskCmd.MarkFlagRequired("wireless-device-id")
	iotwirelessCmd.AddCommand(iotwireless_disassociateWirelessDeviceFromFuotaTaskCmd)
}
