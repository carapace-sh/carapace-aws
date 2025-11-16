package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tnb_getSolFunctionPackageContentCmd = &cobra.Command{
	Use:   "get-sol-function-package-content",
	Short: "Gets the contents of a function package.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tnb_getSolFunctionPackageContentCmd).Standalone()

	tnb_getSolFunctionPackageContentCmd.Flags().String("accept", "", "The format of the package that you want to download from the function packages.")
	tnb_getSolFunctionPackageContentCmd.Flags().String("vnf-pkg-id", "", "ID of the function package.")
	tnb_getSolFunctionPackageContentCmd.MarkFlagRequired("accept")
	tnb_getSolFunctionPackageContentCmd.MarkFlagRequired("vnf-pkg-id")
	tnbCmd.AddCommand(tnb_getSolFunctionPackageContentCmd)
}
