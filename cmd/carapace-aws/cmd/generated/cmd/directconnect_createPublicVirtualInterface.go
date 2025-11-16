package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directconnect_createPublicVirtualInterfaceCmd = &cobra.Command{
	Use:   "create-public-virtual-interface",
	Short: "Creates a public virtual interface.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directconnect_createPublicVirtualInterfaceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(directconnect_createPublicVirtualInterfaceCmd).Standalone()

		directconnect_createPublicVirtualInterfaceCmd.Flags().String("connection-id", "", "The ID of the connection.")
		directconnect_createPublicVirtualInterfaceCmd.Flags().String("new-public-virtual-interface", "", "Information about the public virtual interface.")
		directconnect_createPublicVirtualInterfaceCmd.MarkFlagRequired("connection-id")
		directconnect_createPublicVirtualInterfaceCmd.MarkFlagRequired("new-public-virtual-interface")
	})
	directconnectCmd.AddCommand(directconnect_createPublicVirtualInterfaceCmd)
}
