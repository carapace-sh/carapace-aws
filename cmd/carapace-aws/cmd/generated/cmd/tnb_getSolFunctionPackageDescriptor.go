package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tnb_getSolFunctionPackageDescriptorCmd = &cobra.Command{
	Use:   "get-sol-function-package-descriptor",
	Short: "Gets a function package descriptor in a function package.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tnb_getSolFunctionPackageDescriptorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(tnb_getSolFunctionPackageDescriptorCmd).Standalone()

		tnb_getSolFunctionPackageDescriptorCmd.Flags().String("accept", "", "Indicates which content types, expressed as MIME types, the client is able to understand.")
		tnb_getSolFunctionPackageDescriptorCmd.Flags().String("vnf-pkg-id", "", "ID of the function package.")
		tnb_getSolFunctionPackageDescriptorCmd.MarkFlagRequired("accept")
		tnb_getSolFunctionPackageDescriptorCmd.MarkFlagRequired("vnf-pkg-id")
	})
	tnbCmd.AddCommand(tnb_getSolFunctionPackageDescriptorCmd)
}
