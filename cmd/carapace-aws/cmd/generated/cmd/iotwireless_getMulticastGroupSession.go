package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_getMulticastGroupSessionCmd = &cobra.Command{
	Use:   "get-multicast-group-session",
	Short: "Gets information about a multicast group session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_getMulticastGroupSessionCmd).Standalone()

	iotwireless_getMulticastGroupSessionCmd.Flags().String("id", "", "")
	iotwireless_getMulticastGroupSessionCmd.MarkFlagRequired("id")
	iotwirelessCmd.AddCommand(iotwireless_getMulticastGroupSessionCmd)
}
