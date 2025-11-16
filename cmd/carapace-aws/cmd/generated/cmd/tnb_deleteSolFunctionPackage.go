package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tnb_deleteSolFunctionPackageCmd = &cobra.Command{
	Use:   "delete-sol-function-package",
	Short: "Deletes a function package.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tnb_deleteSolFunctionPackageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(tnb_deleteSolFunctionPackageCmd).Standalone()

		tnb_deleteSolFunctionPackageCmd.Flags().String("vnf-pkg-id", "", "ID of the function package.")
		tnb_deleteSolFunctionPackageCmd.MarkFlagRequired("vnf-pkg-id")
	})
	tnbCmd.AddCommand(tnb_deleteSolFunctionPackageCmd)
}
