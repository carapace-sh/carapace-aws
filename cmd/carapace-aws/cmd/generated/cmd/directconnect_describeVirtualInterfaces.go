package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directconnect_describeVirtualInterfacesCmd = &cobra.Command{
	Use:   "describe-virtual-interfaces",
	Short: "Displays all virtual interfaces for an Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directconnect_describeVirtualInterfacesCmd).Standalone()

	directconnect_describeVirtualInterfacesCmd.Flags().String("connection-id", "", "The ID of the connection.")
	directconnect_describeVirtualInterfacesCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
	directconnect_describeVirtualInterfacesCmd.Flags().String("next-token", "", "The token for the next page of results.")
	directconnect_describeVirtualInterfacesCmd.Flags().String("virtual-interface-id", "", "The ID of the virtual interface.")
	directconnectCmd.AddCommand(directconnect_describeVirtualInterfacesCmd)
}
