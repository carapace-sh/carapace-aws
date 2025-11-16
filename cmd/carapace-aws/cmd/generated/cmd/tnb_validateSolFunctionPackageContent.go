package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tnb_validateSolFunctionPackageContentCmd = &cobra.Command{
	Use:   "validate-sol-function-package-content",
	Short: "Validates function package content.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tnb_validateSolFunctionPackageContentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(tnb_validateSolFunctionPackageContentCmd).Standalone()

		tnb_validateSolFunctionPackageContentCmd.Flags().String("content-type", "", "Function package content type.")
		tnb_validateSolFunctionPackageContentCmd.Flags().String("file", "", "Function package file.")
		tnb_validateSolFunctionPackageContentCmd.Flags().String("vnf-pkg-id", "", "Function package ID.")
		tnb_validateSolFunctionPackageContentCmd.MarkFlagRequired("file")
		tnb_validateSolFunctionPackageContentCmd.MarkFlagRequired("vnf-pkg-id")
	})
	tnbCmd.AddCommand(tnb_validateSolFunctionPackageContentCmd)
}
