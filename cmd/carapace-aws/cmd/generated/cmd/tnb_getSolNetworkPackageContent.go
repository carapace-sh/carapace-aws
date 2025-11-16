package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tnb_getSolNetworkPackageContentCmd = &cobra.Command{
	Use:   "get-sol-network-package-content",
	Short: "Gets the contents of a network package.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tnb_getSolNetworkPackageContentCmd).Standalone()

	tnb_getSolNetworkPackageContentCmd.Flags().String("accept", "", "The format of the package you want to download from the network package.")
	tnb_getSolNetworkPackageContentCmd.Flags().String("nsd-info-id", "", "ID of the network service descriptor in the network package.")
	tnb_getSolNetworkPackageContentCmd.MarkFlagRequired("accept")
	tnb_getSolNetworkPackageContentCmd.MarkFlagRequired("nsd-info-id")
	tnbCmd.AddCommand(tnb_getSolNetworkPackageContentCmd)
}
