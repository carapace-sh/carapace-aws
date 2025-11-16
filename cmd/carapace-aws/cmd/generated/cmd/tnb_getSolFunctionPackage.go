package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tnb_getSolFunctionPackageCmd = &cobra.Command{
	Use:   "get-sol-function-package",
	Short: "Gets the details of an individual function package, such as the operational state and whether the package is in use.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tnb_getSolFunctionPackageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(tnb_getSolFunctionPackageCmd).Standalone()

		tnb_getSolFunctionPackageCmd.Flags().String("vnf-pkg-id", "", "ID of the function package.")
		tnb_getSolFunctionPackageCmd.MarkFlagRequired("vnf-pkg-id")
	})
	tnbCmd.AddCommand(tnb_getSolFunctionPackageCmd)
}
