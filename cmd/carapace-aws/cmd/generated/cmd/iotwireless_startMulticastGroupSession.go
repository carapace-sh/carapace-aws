package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_startMulticastGroupSessionCmd = &cobra.Command{
	Use:   "start-multicast-group-session",
	Short: "Starts a multicast group session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_startMulticastGroupSessionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotwireless_startMulticastGroupSessionCmd).Standalone()

		iotwireless_startMulticastGroupSessionCmd.Flags().String("id", "", "")
		iotwireless_startMulticastGroupSessionCmd.Flags().String("lo-ra-wan", "", "")
		iotwireless_startMulticastGroupSessionCmd.MarkFlagRequired("id")
		iotwireless_startMulticastGroupSessionCmd.MarkFlagRequired("lo-ra-wan")
	})
	iotwirelessCmd.AddCommand(iotwireless_startMulticastGroupSessionCmd)
}
