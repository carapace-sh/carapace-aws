package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_createFuotaTaskCmd = &cobra.Command{
	Use:   "create-fuota-task",
	Short: "Creates a FUOTA task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_createFuotaTaskCmd).Standalone()

	iotwireless_createFuotaTaskCmd.Flags().String("client-request-token", "", "")
	iotwireless_createFuotaTaskCmd.Flags().String("description", "", "")
	iotwireless_createFuotaTaskCmd.Flags().String("descriptor", "", "")
	iotwireless_createFuotaTaskCmd.Flags().String("firmware-update-image", "", "")
	iotwireless_createFuotaTaskCmd.Flags().String("firmware-update-role", "", "")
	iotwireless_createFuotaTaskCmd.Flags().String("fragment-interval-ms", "", "")
	iotwireless_createFuotaTaskCmd.Flags().String("fragment-size-bytes", "", "")
	iotwireless_createFuotaTaskCmd.Flags().String("lo-ra-wan", "", "")
	iotwireless_createFuotaTaskCmd.Flags().String("name", "", "")
	iotwireless_createFuotaTaskCmd.Flags().String("redundancy-percent", "", "")
	iotwireless_createFuotaTaskCmd.Flags().String("tags", "", "")
	iotwireless_createFuotaTaskCmd.MarkFlagRequired("firmware-update-image")
	iotwireless_createFuotaTaskCmd.MarkFlagRequired("firmware-update-role")
	iotwirelessCmd.AddCommand(iotwireless_createFuotaTaskCmd)
}
