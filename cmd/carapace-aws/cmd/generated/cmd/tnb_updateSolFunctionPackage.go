package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tnb_updateSolFunctionPackageCmd = &cobra.Command{
	Use:   "update-sol-function-package",
	Short: "Updates the operational state of function package.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tnb_updateSolFunctionPackageCmd).Standalone()

	tnb_updateSolFunctionPackageCmd.Flags().String("operational-state", "", "Operational state of the function package.")
	tnb_updateSolFunctionPackageCmd.Flags().String("vnf-pkg-id", "", "ID of the function package.")
	tnb_updateSolFunctionPackageCmd.MarkFlagRequired("operational-state")
	tnb_updateSolFunctionPackageCmd.MarkFlagRequired("vnf-pkg-id")
	tnbCmd.AddCommand(tnb_updateSolFunctionPackageCmd)
}
