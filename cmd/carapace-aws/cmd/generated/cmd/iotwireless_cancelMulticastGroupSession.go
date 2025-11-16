package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_cancelMulticastGroupSessionCmd = &cobra.Command{
	Use:   "cancel-multicast-group-session",
	Short: "Cancels an existing multicast group session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_cancelMulticastGroupSessionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotwireless_cancelMulticastGroupSessionCmd).Standalone()

		iotwireless_cancelMulticastGroupSessionCmd.Flags().String("id", "", "")
		iotwireless_cancelMulticastGroupSessionCmd.MarkFlagRequired("id")
	})
	iotwirelessCmd.AddCommand(iotwireless_cancelMulticastGroupSessionCmd)
}
