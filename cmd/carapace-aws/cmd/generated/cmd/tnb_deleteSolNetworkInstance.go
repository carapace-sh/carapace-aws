package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tnb_deleteSolNetworkInstanceCmd = &cobra.Command{
	Use:   "delete-sol-network-instance",
	Short: "Deletes a network instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tnb_deleteSolNetworkInstanceCmd).Standalone()

	tnb_deleteSolNetworkInstanceCmd.Flags().String("ns-instance-id", "", "Network instance ID.")
	tnb_deleteSolNetworkInstanceCmd.MarkFlagRequired("ns-instance-id")
	tnbCmd.AddCommand(tnb_deleteSolNetworkInstanceCmd)
}
