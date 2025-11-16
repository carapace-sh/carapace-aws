package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_updateMulticastGroupCmd = &cobra.Command{
	Use:   "update-multicast-group",
	Short: "Updates properties of a multicast group session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_updateMulticastGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotwireless_updateMulticastGroupCmd).Standalone()

		iotwireless_updateMulticastGroupCmd.Flags().String("description", "", "")
		iotwireless_updateMulticastGroupCmd.Flags().String("id", "", "")
		iotwireless_updateMulticastGroupCmd.Flags().String("lo-ra-wan", "", "")
		iotwireless_updateMulticastGroupCmd.Flags().String("name", "", "")
		iotwireless_updateMulticastGroupCmd.MarkFlagRequired("id")
	})
	iotwirelessCmd.AddCommand(iotwireless_updateMulticastGroupCmd)
}
