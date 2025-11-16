package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_startFuotaTaskCmd = &cobra.Command{
	Use:   "start-fuota-task",
	Short: "Starts a FUOTA task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_startFuotaTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotwireless_startFuotaTaskCmd).Standalone()

		iotwireless_startFuotaTaskCmd.Flags().String("id", "", "")
		iotwireless_startFuotaTaskCmd.Flags().String("lo-ra-wan", "", "")
		iotwireless_startFuotaTaskCmd.MarkFlagRequired("id")
	})
	iotwirelessCmd.AddCommand(iotwireless_startFuotaTaskCmd)
}
