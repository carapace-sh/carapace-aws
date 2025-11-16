package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_getPositionCmd = &cobra.Command{
	Use:   "get-position",
	Short: "Get the position information for a given resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_getPositionCmd).Standalone()

	iotwireless_getPositionCmd.Flags().String("resource-identifier", "", "Resource identifier used to retrieve the position information.")
	iotwireless_getPositionCmd.Flags().String("resource-type", "", "Resource type of the resource for which position information is retrieved.")
	iotwireless_getPositionCmd.MarkFlagRequired("resource-identifier")
	iotwireless_getPositionCmd.MarkFlagRequired("resource-type")
	iotwirelessCmd.AddCommand(iotwireless_getPositionCmd)
}
