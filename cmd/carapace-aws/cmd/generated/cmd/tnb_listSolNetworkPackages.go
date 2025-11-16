package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tnb_listSolNetworkPackagesCmd = &cobra.Command{
	Use:   "list-sol-network-packages",
	Short: "Lists network packages.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tnb_listSolNetworkPackagesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(tnb_listSolNetworkPackagesCmd).Standalone()

		tnb_listSolNetworkPackagesCmd.Flags().String("max-results", "", "The maximum number of results to include in the response.")
		tnb_listSolNetworkPackagesCmd.Flags().String("next-token", "", "The token for the next page of results.")
	})
	tnbCmd.AddCommand(tnb_listSolNetworkPackagesCmd)
}
