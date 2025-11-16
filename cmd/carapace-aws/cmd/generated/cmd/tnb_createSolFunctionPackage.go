package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tnb_createSolFunctionPackageCmd = &cobra.Command{
	Use:   "create-sol-function-package",
	Short: "Creates a function package.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tnb_createSolFunctionPackageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(tnb_createSolFunctionPackageCmd).Standalone()

		tnb_createSolFunctionPackageCmd.Flags().String("tags", "", "A tag is a label that you assign to an Amazon Web Services resource.")
	})
	tnbCmd.AddCommand(tnb_createSolFunctionPackageCmd)
}
