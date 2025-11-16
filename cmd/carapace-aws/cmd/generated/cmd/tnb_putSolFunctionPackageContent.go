package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tnb_putSolFunctionPackageContentCmd = &cobra.Command{
	Use:   "put-sol-function-package-content",
	Short: "Uploads the contents of a function package.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tnb_putSolFunctionPackageContentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(tnb_putSolFunctionPackageContentCmd).Standalone()

		tnb_putSolFunctionPackageContentCmd.Flags().String("content-type", "", "Function package content type.")
		tnb_putSolFunctionPackageContentCmd.Flags().String("file", "", "Function package file.")
		tnb_putSolFunctionPackageContentCmd.Flags().String("vnf-pkg-id", "", "Function package ID.")
		tnb_putSolFunctionPackageContentCmd.MarkFlagRequired("file")
		tnb_putSolFunctionPackageContentCmd.MarkFlagRequired("vnf-pkg-id")
	})
	tnbCmd.AddCommand(tnb_putSolFunctionPackageContentCmd)
}
