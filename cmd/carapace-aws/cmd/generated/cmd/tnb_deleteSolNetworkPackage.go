package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tnb_deleteSolNetworkPackageCmd = &cobra.Command{
	Use:   "delete-sol-network-package",
	Short: "Deletes network package.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tnb_deleteSolNetworkPackageCmd).Standalone()

	tnb_deleteSolNetworkPackageCmd.Flags().String("nsd-info-id", "", "ID of the network service descriptor in the network package.")
	tnb_deleteSolNetworkPackageCmd.MarkFlagRequired("nsd-info-id")
	tnbCmd.AddCommand(tnb_deleteSolNetworkPackageCmd)
}
