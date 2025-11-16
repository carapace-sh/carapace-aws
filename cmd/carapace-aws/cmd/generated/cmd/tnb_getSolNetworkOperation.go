package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tnb_getSolNetworkOperationCmd = &cobra.Command{
	Use:   "get-sol-network-operation",
	Short: "Gets the details of a network operation, including the tasks involved in the network operation and the status of the tasks.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tnb_getSolNetworkOperationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(tnb_getSolNetworkOperationCmd).Standalone()

		tnb_getSolNetworkOperationCmd.Flags().String("ns-lcm-op-occ-id", "", "The identifier of the network operation.")
		tnb_getSolNetworkOperationCmd.MarkFlagRequired("ns-lcm-op-occ-id")
	})
	tnbCmd.AddCommand(tnb_getSolNetworkOperationCmd)
}
