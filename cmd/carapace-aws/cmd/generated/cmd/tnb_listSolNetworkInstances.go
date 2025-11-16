package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tnb_listSolNetworkInstancesCmd = &cobra.Command{
	Use:   "list-sol-network-instances",
	Short: "Lists your network instances.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tnb_listSolNetworkInstancesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(tnb_listSolNetworkInstancesCmd).Standalone()

		tnb_listSolNetworkInstancesCmd.Flags().String("max-results", "", "The maximum number of results to include in the response.")
		tnb_listSolNetworkInstancesCmd.Flags().String("next-token", "", "The token for the next page of results.")
	})
	tnbCmd.AddCommand(tnb_listSolNetworkInstancesCmd)
}
