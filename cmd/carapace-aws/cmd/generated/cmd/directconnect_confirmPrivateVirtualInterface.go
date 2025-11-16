package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directconnect_confirmPrivateVirtualInterfaceCmd = &cobra.Command{
	Use:   "confirm-private-virtual-interface",
	Short: "Accepts ownership of a private virtual interface created by another Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directconnect_confirmPrivateVirtualInterfaceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(directconnect_confirmPrivateVirtualInterfaceCmd).Standalone()

		directconnect_confirmPrivateVirtualInterfaceCmd.Flags().String("direct-connect-gateway-id", "", "The ID of the Direct Connect gateway.")
		directconnect_confirmPrivateVirtualInterfaceCmd.Flags().String("virtual-gateway-id", "", "The ID of the virtual private gateway.")
		directconnect_confirmPrivateVirtualInterfaceCmd.Flags().String("virtual-interface-id", "", "The ID of the virtual interface.")
		directconnect_confirmPrivateVirtualInterfaceCmd.MarkFlagRequired("virtual-interface-id")
	})
	directconnectCmd.AddCommand(directconnect_confirmPrivateVirtualInterfaceCmd)
}
