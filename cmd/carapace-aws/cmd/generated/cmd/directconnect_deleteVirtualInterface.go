package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directconnect_deleteVirtualInterfaceCmd = &cobra.Command{
	Use:   "delete-virtual-interface",
	Short: "Deletes a virtual interface.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directconnect_deleteVirtualInterfaceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(directconnect_deleteVirtualInterfaceCmd).Standalone()

		directconnect_deleteVirtualInterfaceCmd.Flags().String("virtual-interface-id", "", "The ID of the virtual interface.")
		directconnect_deleteVirtualInterfaceCmd.MarkFlagRequired("virtual-interface-id")
	})
	directconnectCmd.AddCommand(directconnect_deleteVirtualInterfaceCmd)
}
