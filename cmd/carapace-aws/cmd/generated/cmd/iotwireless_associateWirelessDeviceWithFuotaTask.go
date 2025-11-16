package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_associateWirelessDeviceWithFuotaTaskCmd = &cobra.Command{
	Use:   "associate-wireless-device-with-fuota-task",
	Short: "Associate a wireless device with a FUOTA task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_associateWirelessDeviceWithFuotaTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotwireless_associateWirelessDeviceWithFuotaTaskCmd).Standalone()

		iotwireless_associateWirelessDeviceWithFuotaTaskCmd.Flags().String("id", "", "")
		iotwireless_associateWirelessDeviceWithFuotaTaskCmd.Flags().String("wireless-device-id", "", "")
		iotwireless_associateWirelessDeviceWithFuotaTaskCmd.MarkFlagRequired("id")
		iotwireless_associateWirelessDeviceWithFuotaTaskCmd.MarkFlagRequired("wireless-device-id")
	})
	iotwirelessCmd.AddCommand(iotwireless_associateWirelessDeviceWithFuotaTaskCmd)
}
