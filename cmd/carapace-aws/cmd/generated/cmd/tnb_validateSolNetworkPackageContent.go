package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tnb_validateSolNetworkPackageContentCmd = &cobra.Command{
	Use:   "validate-sol-network-package-content",
	Short: "Validates network package content.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tnb_validateSolNetworkPackageContentCmd).Standalone()

	tnb_validateSolNetworkPackageContentCmd.Flags().String("content-type", "", "Network package content type.")
	tnb_validateSolNetworkPackageContentCmd.Flags().String("file", "", "Network package file.")
	tnb_validateSolNetworkPackageContentCmd.Flags().String("nsd-info-id", "", "Network service descriptor file.")
	tnb_validateSolNetworkPackageContentCmd.MarkFlagRequired("file")
	tnb_validateSolNetworkPackageContentCmd.MarkFlagRequired("nsd-info-id")
	tnbCmd.AddCommand(tnb_validateSolNetworkPackageContentCmd)
}
