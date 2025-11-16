package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_createMulticastGroupCmd = &cobra.Command{
	Use:   "create-multicast-group",
	Short: "Creates a multicast group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_createMulticastGroupCmd).Standalone()

	iotwireless_createMulticastGroupCmd.Flags().String("client-request-token", "", "Each resource must have a unique client request token.")
	iotwireless_createMulticastGroupCmd.Flags().String("description", "", "The description of the multicast group.")
	iotwireless_createMulticastGroupCmd.Flags().String("lo-ra-wan", "", "")
	iotwireless_createMulticastGroupCmd.Flags().String("name", "", "")
	iotwireless_createMulticastGroupCmd.Flags().String("tags", "", "")
	iotwireless_createMulticastGroupCmd.MarkFlagRequired("lo-ra-wan")
	iotwirelessCmd.AddCommand(iotwireless_createMulticastGroupCmd)
}
