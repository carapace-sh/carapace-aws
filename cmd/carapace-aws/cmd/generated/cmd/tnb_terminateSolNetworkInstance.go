package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tnb_terminateSolNetworkInstanceCmd = &cobra.Command{
	Use:   "terminate-sol-network-instance",
	Short: "Terminates a network instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tnb_terminateSolNetworkInstanceCmd).Standalone()

	tnb_terminateSolNetworkInstanceCmd.Flags().String("ns-instance-id", "", "ID of the network instance.")
	tnb_terminateSolNetworkInstanceCmd.Flags().String("tags", "", "A tag is a label that you assign to an Amazon Web Services resource.")
	tnb_terminateSolNetworkInstanceCmd.MarkFlagRequired("ns-instance-id")
	tnbCmd.AddCommand(tnb_terminateSolNetworkInstanceCmd)
}
