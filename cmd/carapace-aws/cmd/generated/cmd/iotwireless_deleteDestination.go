package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_deleteDestinationCmd = &cobra.Command{
	Use:   "delete-destination",
	Short: "Deletes a destination.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_deleteDestinationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotwireless_deleteDestinationCmd).Standalone()

		iotwireless_deleteDestinationCmd.Flags().String("name", "", "The name of the resource to delete.")
		iotwireless_deleteDestinationCmd.MarkFlagRequired("name")
	})
	iotwirelessCmd.AddCommand(iotwireless_deleteDestinationCmd)
}
