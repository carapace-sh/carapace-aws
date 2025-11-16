package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directconnect_allocatePublicVirtualInterfaceCmd = &cobra.Command{
	Use:   "allocate-public-virtual-interface",
	Short: "Provisions a public virtual interface to be owned by the specified Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directconnect_allocatePublicVirtualInterfaceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(directconnect_allocatePublicVirtualInterfaceCmd).Standalone()

		directconnect_allocatePublicVirtualInterfaceCmd.Flags().String("connection-id", "", "The ID of the connection on which the public virtual interface is provisioned.")
		directconnect_allocatePublicVirtualInterfaceCmd.Flags().String("new-public-virtual-interface-allocation", "", "Information about the public virtual interface.")
		directconnect_allocatePublicVirtualInterfaceCmd.Flags().String("owner-account", "", "The ID of the Amazon Web Services account that owns the public virtual interface.")
		directconnect_allocatePublicVirtualInterfaceCmd.MarkFlagRequired("connection-id")
		directconnect_allocatePublicVirtualInterfaceCmd.MarkFlagRequired("new-public-virtual-interface-allocation")
		directconnect_allocatePublicVirtualInterfaceCmd.MarkFlagRequired("owner-account")
	})
	directconnectCmd.AddCommand(directconnect_allocatePublicVirtualInterfaceCmd)
}
