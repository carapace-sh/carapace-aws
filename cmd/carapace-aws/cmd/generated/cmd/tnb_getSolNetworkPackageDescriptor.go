package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tnb_getSolNetworkPackageDescriptorCmd = &cobra.Command{
	Use:   "get-sol-network-package-descriptor",
	Short: "Gets the content of the network service descriptor.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tnb_getSolNetworkPackageDescriptorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(tnb_getSolNetworkPackageDescriptorCmd).Standalone()

		tnb_getSolNetworkPackageDescriptorCmd.Flags().String("nsd-info-id", "", "ID of the network service descriptor in the network package.")
		tnb_getSolNetworkPackageDescriptorCmd.MarkFlagRequired("nsd-info-id")
	})
	tnbCmd.AddCommand(tnb_getSolNetworkPackageDescriptorCmd)
}
