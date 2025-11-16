package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directconnect_allocatePrivateVirtualInterfaceCmd = &cobra.Command{
	Use:   "allocate-private-virtual-interface",
	Short: "Provisions a private virtual interface to be owned by the specified Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directconnect_allocatePrivateVirtualInterfaceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(directconnect_allocatePrivateVirtualInterfaceCmd).Standalone()

		directconnect_allocatePrivateVirtualInterfaceCmd.Flags().String("connection-id", "", "The ID of the connection on which the private virtual interface is provisioned.")
		directconnect_allocatePrivateVirtualInterfaceCmd.Flags().String("new-private-virtual-interface-allocation", "", "Information about the private virtual interface.")
		directconnect_allocatePrivateVirtualInterfaceCmd.Flags().String("owner-account", "", "The ID of the Amazon Web Services account that owns the virtual private interface.")
		directconnect_allocatePrivateVirtualInterfaceCmd.MarkFlagRequired("connection-id")
		directconnect_allocatePrivateVirtualInterfaceCmd.MarkFlagRequired("new-private-virtual-interface-allocation")
		directconnect_allocatePrivateVirtualInterfaceCmd.MarkFlagRequired("owner-account")
	})
	directconnectCmd.AddCommand(directconnect_allocatePrivateVirtualInterfaceCmd)
}
