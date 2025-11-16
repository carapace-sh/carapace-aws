package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_updatePositionCmd = &cobra.Command{
	Use:   "update-position",
	Short: "Update the position information of a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_updatePositionCmd).Standalone()

	iotwireless_updatePositionCmd.Flags().String("position", "", "The position information of the resource.")
	iotwireless_updatePositionCmd.Flags().String("resource-identifier", "", "Resource identifier of the resource for which position is updated.")
	iotwireless_updatePositionCmd.Flags().String("resource-type", "", "Resource type of the resource for which position is updated.")
	iotwireless_updatePositionCmd.MarkFlagRequired("position")
	iotwireless_updatePositionCmd.MarkFlagRequired("resource-identifier")
	iotwireless_updatePositionCmd.MarkFlagRequired("resource-type")
	iotwirelessCmd.AddCommand(iotwireless_updatePositionCmd)
}
