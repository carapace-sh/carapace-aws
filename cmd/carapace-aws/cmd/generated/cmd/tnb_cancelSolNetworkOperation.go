package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tnb_cancelSolNetworkOperationCmd = &cobra.Command{
	Use:   "cancel-sol-network-operation",
	Short: "Cancels a network operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tnb_cancelSolNetworkOperationCmd).Standalone()

	tnb_cancelSolNetworkOperationCmd.Flags().String("ns-lcm-op-occ-id", "", "The identifier of the network operation.")
	tnb_cancelSolNetworkOperationCmd.MarkFlagRequired("ns-lcm-op-occ-id")
	tnbCmd.AddCommand(tnb_cancelSolNetworkOperationCmd)
}
