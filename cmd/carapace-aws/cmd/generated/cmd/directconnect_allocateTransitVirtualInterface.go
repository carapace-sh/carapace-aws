package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directconnect_allocateTransitVirtualInterfaceCmd = &cobra.Command{
	Use:   "allocate-transit-virtual-interface",
	Short: "Provisions a transit virtual interface to be owned by the specified Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directconnect_allocateTransitVirtualInterfaceCmd).Standalone()

	directconnect_allocateTransitVirtualInterfaceCmd.Flags().String("connection-id", "", "The ID of the connection on which the transit virtual interface is provisioned.")
	directconnect_allocateTransitVirtualInterfaceCmd.Flags().String("new-transit-virtual-interface-allocation", "", "Information about the transit virtual interface.")
	directconnect_allocateTransitVirtualInterfaceCmd.Flags().String("owner-account", "", "The ID of the Amazon Web Services account that owns the transit virtual interface.")
	directconnect_allocateTransitVirtualInterfaceCmd.MarkFlagRequired("connection-id")
	directconnect_allocateTransitVirtualInterfaceCmd.MarkFlagRequired("new-transit-virtual-interface-allocation")
	directconnect_allocateTransitVirtualInterfaceCmd.MarkFlagRequired("owner-account")
	directconnectCmd.AddCommand(directconnect_allocateTransitVirtualInterfaceCmd)
}
