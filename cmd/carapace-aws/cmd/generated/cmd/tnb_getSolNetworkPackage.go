package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tnb_getSolNetworkPackageCmd = &cobra.Command{
	Use:   "get-sol-network-package",
	Short: "Gets the details of a network package.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tnb_getSolNetworkPackageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(tnb_getSolNetworkPackageCmd).Standalone()

		tnb_getSolNetworkPackageCmd.Flags().String("nsd-info-id", "", "ID of the network service descriptor in the network package.")
		tnb_getSolNetworkPackageCmd.MarkFlagRequired("nsd-info-id")
	})
	tnbCmd.AddCommand(tnb_getSolNetworkPackageCmd)
}
