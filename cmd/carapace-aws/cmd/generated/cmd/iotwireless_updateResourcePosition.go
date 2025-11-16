package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_updateResourcePositionCmd = &cobra.Command{
	Use:   "update-resource-position",
	Short: "Update the position information of a given wireless device or a wireless gateway resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_updateResourcePositionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotwireless_updateResourcePositionCmd).Standalone()

		iotwireless_updateResourcePositionCmd.Flags().String("geo-json-payload", "", "The position information of the resource, displayed as a JSON payload.")
		iotwireless_updateResourcePositionCmd.Flags().String("resource-identifier", "", "The identifier of the resource for which position information is updated.")
		iotwireless_updateResourcePositionCmd.Flags().String("resource-type", "", "The type of resource for which position information is updated, which can be a wireless device or a wireless gateway.")
		iotwireless_updateResourcePositionCmd.MarkFlagRequired("resource-identifier")
		iotwireless_updateResourcePositionCmd.MarkFlagRequired("resource-type")
	})
	iotwirelessCmd.AddCommand(iotwireless_updateResourcePositionCmd)
}
