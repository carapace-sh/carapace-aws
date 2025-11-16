package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_getResourcePositionCmd = &cobra.Command{
	Use:   "get-resource-position",
	Short: "Get the position information for a given wireless device or a wireless gateway resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_getResourcePositionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotwireless_getResourcePositionCmd).Standalone()

		iotwireless_getResourcePositionCmd.Flags().String("resource-identifier", "", "The identifier of the resource for which position information is retrieved.")
		iotwireless_getResourcePositionCmd.Flags().String("resource-type", "", "The type of resource for which position information is retrieved, which can be a wireless device or a wireless gateway.")
		iotwireless_getResourcePositionCmd.MarkFlagRequired("resource-identifier")
		iotwireless_getResourcePositionCmd.MarkFlagRequired("resource-type")
	})
	iotwirelessCmd.AddCommand(iotwireless_getResourcePositionCmd)
}
