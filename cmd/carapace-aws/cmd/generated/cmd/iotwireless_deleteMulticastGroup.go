package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_deleteMulticastGroupCmd = &cobra.Command{
	Use:   "delete-multicast-group",
	Short: "Deletes a multicast group if it is not in use by a FUOTA task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_deleteMulticastGroupCmd).Standalone()

	iotwireless_deleteMulticastGroupCmd.Flags().String("id", "", "")
	iotwireless_deleteMulticastGroupCmd.MarkFlagRequired("id")
	iotwirelessCmd.AddCommand(iotwireless_deleteMulticastGroupCmd)
}
