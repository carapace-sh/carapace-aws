package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_getDestinationCmd = &cobra.Command{
	Use:   "get-destination",
	Short: "Gets information about a destination.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_getDestinationCmd).Standalone()

	iotwireless_getDestinationCmd.Flags().String("name", "", "The name of the resource to get.")
	iotwireless_getDestinationCmd.MarkFlagRequired("name")
	iotwirelessCmd.AddCommand(iotwireless_getDestinationCmd)
}
