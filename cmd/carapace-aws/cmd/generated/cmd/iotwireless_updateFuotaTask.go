package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_updateFuotaTaskCmd = &cobra.Command{
	Use:   "update-fuota-task",
	Short: "Updates properties of a FUOTA task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_updateFuotaTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotwireless_updateFuotaTaskCmd).Standalone()

		iotwireless_updateFuotaTaskCmd.Flags().String("description", "", "")
		iotwireless_updateFuotaTaskCmd.Flags().String("descriptor", "", "")
		iotwireless_updateFuotaTaskCmd.Flags().String("firmware-update-image", "", "")
		iotwireless_updateFuotaTaskCmd.Flags().String("firmware-update-role", "", "")
		iotwireless_updateFuotaTaskCmd.Flags().String("fragment-interval-ms", "", "")
		iotwireless_updateFuotaTaskCmd.Flags().String("fragment-size-bytes", "", "")
		iotwireless_updateFuotaTaskCmd.Flags().String("id", "", "")
		iotwireless_updateFuotaTaskCmd.Flags().String("lo-ra-wan", "", "")
		iotwireless_updateFuotaTaskCmd.Flags().String("name", "", "")
		iotwireless_updateFuotaTaskCmd.Flags().String("redundancy-percent", "", "")
		iotwireless_updateFuotaTaskCmd.MarkFlagRequired("id")
	})
	iotwirelessCmd.AddCommand(iotwireless_updateFuotaTaskCmd)
}
