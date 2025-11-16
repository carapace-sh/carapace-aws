package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directconnect_createPrivateVirtualInterfaceCmd = &cobra.Command{
	Use:   "create-private-virtual-interface",
	Short: "Creates a private virtual interface.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directconnect_createPrivateVirtualInterfaceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(directconnect_createPrivateVirtualInterfaceCmd).Standalone()

		directconnect_createPrivateVirtualInterfaceCmd.Flags().String("connection-id", "", "The ID of the connection.")
		directconnect_createPrivateVirtualInterfaceCmd.Flags().String("new-private-virtual-interface", "", "Information about the private virtual interface.")
		directconnect_createPrivateVirtualInterfaceCmd.MarkFlagRequired("connection-id")
		directconnect_createPrivateVirtualInterfaceCmd.MarkFlagRequired("new-private-virtual-interface")
	})
	directconnectCmd.AddCommand(directconnect_createPrivateVirtualInterfaceCmd)
}
