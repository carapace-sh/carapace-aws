package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directconnect_associateVirtualInterfaceCmd = &cobra.Command{
	Use:   "associate-virtual-interface",
	Short: "Associates a virtual interface with a specified link aggregation group (LAG) or connection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directconnect_associateVirtualInterfaceCmd).Standalone()

	directconnect_associateVirtualInterfaceCmd.Flags().String("connection-id", "", "The ID of the LAG or connection.")
	directconnect_associateVirtualInterfaceCmd.Flags().String("virtual-interface-id", "", "The ID of the virtual interface.")
	directconnect_associateVirtualInterfaceCmd.MarkFlagRequired("connection-id")
	directconnect_associateVirtualInterfaceCmd.MarkFlagRequired("virtual-interface-id")
	directconnectCmd.AddCommand(directconnect_associateVirtualInterfaceCmd)
}
