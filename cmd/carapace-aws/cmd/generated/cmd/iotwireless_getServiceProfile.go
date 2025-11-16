package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_getServiceProfileCmd = &cobra.Command{
	Use:   "get-service-profile",
	Short: "Gets information about a service profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_getServiceProfileCmd).Standalone()

	iotwireless_getServiceProfileCmd.Flags().String("id", "", "The ID of the resource to get.")
	iotwireless_getServiceProfileCmd.MarkFlagRequired("id")
	iotwirelessCmd.AddCommand(iotwireless_getServiceProfileCmd)
}
