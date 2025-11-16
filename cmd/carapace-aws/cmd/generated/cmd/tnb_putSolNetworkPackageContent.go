package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tnb_putSolNetworkPackageContentCmd = &cobra.Command{
	Use:   "put-sol-network-package-content",
	Short: "Uploads the contents of a network package.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tnb_putSolNetworkPackageContentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(tnb_putSolNetworkPackageContentCmd).Standalone()

		tnb_putSolNetworkPackageContentCmd.Flags().String("content-type", "", "Network package content type.")
		tnb_putSolNetworkPackageContentCmd.Flags().String("file", "", "Network package file.")
		tnb_putSolNetworkPackageContentCmd.Flags().String("nsd-info-id", "", "Network service descriptor info ID.")
		tnb_putSolNetworkPackageContentCmd.MarkFlagRequired("file")
		tnb_putSolNetworkPackageContentCmd.MarkFlagRequired("nsd-info-id")
	})
	tnbCmd.AddCommand(tnb_putSolNetworkPackageContentCmd)
}
