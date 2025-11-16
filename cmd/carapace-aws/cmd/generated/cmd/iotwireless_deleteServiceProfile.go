package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_deleteServiceProfileCmd = &cobra.Command{
	Use:   "delete-service-profile",
	Short: "Deletes a service profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_deleteServiceProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotwireless_deleteServiceProfileCmd).Standalone()

		iotwireless_deleteServiceProfileCmd.Flags().String("id", "", "The ID of the resource to delete.")
		iotwireless_deleteServiceProfileCmd.MarkFlagRequired("id")
	})
	iotwirelessCmd.AddCommand(iotwireless_deleteServiceProfileCmd)
}
