package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directconnect_createTransitVirtualInterfaceCmd = &cobra.Command{
	Use:   "create-transit-virtual-interface",
	Short: "Creates a transit virtual interface.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directconnect_createTransitVirtualInterfaceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(directconnect_createTransitVirtualInterfaceCmd).Standalone()

		directconnect_createTransitVirtualInterfaceCmd.Flags().String("connection-id", "", "The ID of the connection.")
		directconnect_createTransitVirtualInterfaceCmd.Flags().String("new-transit-virtual-interface", "", "Information about the transit virtual interface.")
		directconnect_createTransitVirtualInterfaceCmd.MarkFlagRequired("connection-id")
		directconnect_createTransitVirtualInterfaceCmd.MarkFlagRequired("new-transit-virtual-interface")
	})
	directconnectCmd.AddCommand(directconnect_createTransitVirtualInterfaceCmd)
}
