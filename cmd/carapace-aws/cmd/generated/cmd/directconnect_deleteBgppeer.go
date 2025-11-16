package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directconnect_deleteBgppeerCmd = &cobra.Command{
	Use:   "delete-bgppeer",
	Short: "Deletes the specified BGP peer on the specified virtual interface with the specified customer address and ASN.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directconnect_deleteBgppeerCmd).Standalone()

	directconnect_deleteBgppeerCmd.Flags().String("asn", "", "The autonomous system number (ASN).")
	directconnect_deleteBgppeerCmd.Flags().String("asn-long", "", "The long ASN for the BGP peer to be deleted from a Direct Connect virtual interface.")
	directconnect_deleteBgppeerCmd.Flags().String("bgp-peer-id", "", "The ID of the BGP peer.")
	directconnect_deleteBgppeerCmd.Flags().String("customer-address", "", "The IP address assigned to the customer interface.")
	directconnect_deleteBgppeerCmd.Flags().String("virtual-interface-id", "", "The ID of the virtual interface.")
	directconnectCmd.AddCommand(directconnect_deleteBgppeerCmd)
}
