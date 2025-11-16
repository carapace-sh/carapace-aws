package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tnb_listSolFunctionPackagesCmd = &cobra.Command{
	Use:   "list-sol-function-packages",
	Short: "Lists information about function packages.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tnb_listSolFunctionPackagesCmd).Standalone()

	tnb_listSolFunctionPackagesCmd.Flags().String("max-results", "", "The maximum number of results to include in the response.")
	tnb_listSolFunctionPackagesCmd.Flags().String("next-token", "", "The token for the next page of results.")
	tnbCmd.AddCommand(tnb_listSolFunctionPackagesCmd)
}
