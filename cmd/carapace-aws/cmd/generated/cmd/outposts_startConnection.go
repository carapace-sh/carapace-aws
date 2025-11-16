package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var outposts_startConnectionCmd = &cobra.Command{
	Use:   "start-connection",
	Short: "Amazon Web Services uses this action to install Outpost servers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(outposts_startConnectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(outposts_startConnectionCmd).Standalone()

		outposts_startConnectionCmd.Flags().String("asset-id", "", "The ID of the Outpost server.")
		outposts_startConnectionCmd.Flags().String("client-public-key", "", "The public key of the client.")
		outposts_startConnectionCmd.Flags().String("device-serial-number", "", "The serial number of the dongle.")
		outposts_startConnectionCmd.Flags().String("network-interface-device-index", "", "The device index of the network interface on the Outpost server.")
		outposts_startConnectionCmd.MarkFlagRequired("asset-id")
		outposts_startConnectionCmd.MarkFlagRequired("client-public-key")
		outposts_startConnectionCmd.MarkFlagRequired("network-interface-device-index")
	})
	outpostsCmd.AddCommand(outposts_startConnectionCmd)
}
