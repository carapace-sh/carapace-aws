package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tnb_listSolNetworkOperationsCmd = &cobra.Command{
	Use:   "list-sol-network-operations",
	Short: "Lists details for a network operation, including when the operation started and the status of the operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tnb_listSolNetworkOperationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(tnb_listSolNetworkOperationsCmd).Standalone()

		tnb_listSolNetworkOperationsCmd.Flags().String("max-results", "", "The maximum number of results to include in the response.")
		tnb_listSolNetworkOperationsCmd.Flags().String("next-token", "", "The token for the next page of results.")
		tnb_listSolNetworkOperationsCmd.Flags().String("ns-instance-id", "", "Network instance id filter, to retrieve network operations associated to a network instance.")
	})
	tnbCmd.AddCommand(tnb_listSolNetworkOperationsCmd)
}
