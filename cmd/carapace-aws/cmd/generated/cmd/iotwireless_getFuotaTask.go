package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_getFuotaTaskCmd = &cobra.Command{
	Use:   "get-fuota-task",
	Short: "Gets information about a FUOTA task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_getFuotaTaskCmd).Standalone()

	iotwireless_getFuotaTaskCmd.Flags().String("id", "", "")
	iotwireless_getFuotaTaskCmd.MarkFlagRequired("id")
	iotwirelessCmd.AddCommand(iotwireless_getFuotaTaskCmd)
}
