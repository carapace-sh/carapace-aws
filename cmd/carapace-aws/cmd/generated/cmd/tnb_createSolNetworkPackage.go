package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tnb_createSolNetworkPackageCmd = &cobra.Command{
	Use:   "create-sol-network-package",
	Short: "Creates a network package.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tnb_createSolNetworkPackageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(tnb_createSolNetworkPackageCmd).Standalone()

		tnb_createSolNetworkPackageCmd.Flags().String("tags", "", "A tag is a label that you assign to an Amazon Web Services resource.")
	})
	tnbCmd.AddCommand(tnb_createSolNetworkPackageCmd)
}
