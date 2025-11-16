package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directconnect_confirmPublicVirtualInterfaceCmd = &cobra.Command{
	Use:   "confirm-public-virtual-interface",
	Short: "Accepts ownership of a public virtual interface created by another Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directconnect_confirmPublicVirtualInterfaceCmd).Standalone()

	directconnect_confirmPublicVirtualInterfaceCmd.Flags().String("virtual-interface-id", "", "The ID of the virtual interface.")
	directconnect_confirmPublicVirtualInterfaceCmd.MarkFlagRequired("virtual-interface-id")
	directconnectCmd.AddCommand(directconnect_confirmPublicVirtualInterfaceCmd)
}
