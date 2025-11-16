package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tnb_getSolNetworkInstanceCmd = &cobra.Command{
	Use:   "get-sol-network-instance",
	Short: "Gets the details of the network instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tnb_getSolNetworkInstanceCmd).Standalone()

	tnb_getSolNetworkInstanceCmd.Flags().String("ns-instance-id", "", "ID of the network instance.")
	tnb_getSolNetworkInstanceCmd.MarkFlagRequired("ns-instance-id")
	tnbCmd.AddCommand(tnb_getSolNetworkInstanceCmd)
}
