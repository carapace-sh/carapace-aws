package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_getMulticastGroupCmd = &cobra.Command{
	Use:   "get-multicast-group",
	Short: "Gets information about a multicast group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_getMulticastGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotwireless_getMulticastGroupCmd).Standalone()

		iotwireless_getMulticastGroupCmd.Flags().String("id", "", "")
		iotwireless_getMulticastGroupCmd.MarkFlagRequired("id")
	})
	iotwirelessCmd.AddCommand(iotwireless_getMulticastGroupCmd)
}
