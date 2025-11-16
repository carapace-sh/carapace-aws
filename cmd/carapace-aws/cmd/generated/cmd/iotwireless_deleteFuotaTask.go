package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_deleteFuotaTaskCmd = &cobra.Command{
	Use:   "delete-fuota-task",
	Short: "Deletes a FUOTA task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_deleteFuotaTaskCmd).Standalone()

	iotwireless_deleteFuotaTaskCmd.Flags().String("id", "", "")
	iotwireless_deleteFuotaTaskCmd.MarkFlagRequired("id")
	iotwirelessCmd.AddCommand(iotwireless_deleteFuotaTaskCmd)
}
