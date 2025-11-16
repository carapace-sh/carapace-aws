package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tnb_updateSolNetworkPackageCmd = &cobra.Command{
	Use:   "update-sol-network-package",
	Short: "Updates the operational state of a network package.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tnb_updateSolNetworkPackageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(tnb_updateSolNetworkPackageCmd).Standalone()

		tnb_updateSolNetworkPackageCmd.Flags().String("nsd-info-id", "", "ID of the network service descriptor in the network package.")
		tnb_updateSolNetworkPackageCmd.Flags().String("nsd-operational-state", "", "Operational state of the network service descriptor in the network package.")
		tnb_updateSolNetworkPackageCmd.MarkFlagRequired("nsd-info-id")
		tnb_updateSolNetworkPackageCmd.MarkFlagRequired("nsd-operational-state")
	})
	tnbCmd.AddCommand(tnb_updateSolNetworkPackageCmd)
}
