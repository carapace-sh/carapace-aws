package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directconnect_listVirtualInterfaceTestHistoryCmd = &cobra.Command{
	Use:   "list-virtual-interface-test-history",
	Short: "Lists the virtual interface failover test history.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directconnect_listVirtualInterfaceTestHistoryCmd).Standalone()

	directconnect_listVirtualInterfaceTestHistoryCmd.Flags().String("bgp-peers", "", "The BGP peers that were placed in the DOWN state during the virtual interface failover test.")
	directconnect_listVirtualInterfaceTestHistoryCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
	directconnect_listVirtualInterfaceTestHistoryCmd.Flags().String("next-token", "", "The token for the next page of results.")
	directconnect_listVirtualInterfaceTestHistoryCmd.Flags().String("status", "", "The status of the virtual interface failover test.")
	directconnect_listVirtualInterfaceTestHistoryCmd.Flags().String("test-id", "", "The ID of the virtual interface failover test.")
	directconnect_listVirtualInterfaceTestHistoryCmd.Flags().String("virtual-interface-id", "", "The ID of the virtual interface that was tested.")
	directconnectCmd.AddCommand(directconnect_listVirtualInterfaceTestHistoryCmd)
}
