package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directconnect_confirmTransitVirtualInterfaceCmd = &cobra.Command{
	Use:   "confirm-transit-virtual-interface",
	Short: "Accepts ownership of a transit virtual interface created by another Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directconnect_confirmTransitVirtualInterfaceCmd).Standalone()

	directconnect_confirmTransitVirtualInterfaceCmd.Flags().String("direct-connect-gateway-id", "", "The ID of the Direct Connect gateway.")
	directconnect_confirmTransitVirtualInterfaceCmd.Flags().String("virtual-interface-id", "", "The ID of the virtual interface.")
	directconnect_confirmTransitVirtualInterfaceCmd.MarkFlagRequired("direct-connect-gateway-id")
	directconnect_confirmTransitVirtualInterfaceCmd.MarkFlagRequired("virtual-interface-id")
	directconnectCmd.AddCommand(directconnect_confirmTransitVirtualInterfaceCmd)
}
