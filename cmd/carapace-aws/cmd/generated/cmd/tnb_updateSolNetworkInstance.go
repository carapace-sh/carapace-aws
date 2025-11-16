package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tnb_updateSolNetworkInstanceCmd = &cobra.Command{
	Use:   "update-sol-network-instance",
	Short: "Update a network instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tnb_updateSolNetworkInstanceCmd).Standalone()

	tnb_updateSolNetworkInstanceCmd.Flags().String("modify-vnf-info-data", "", "Identifies the network function information parameters and/or the configurable properties of the network function to be modified.")
	tnb_updateSolNetworkInstanceCmd.Flags().String("ns-instance-id", "", "ID of the network instance.")
	tnb_updateSolNetworkInstanceCmd.Flags().String("tags", "", "A tag is a label that you assign to an Amazon Web Services resource.")
	tnb_updateSolNetworkInstanceCmd.Flags().String("update-ns", "", "Identifies the network service descriptor and the configurable properties of the descriptor, to be used for the update.")
	tnb_updateSolNetworkInstanceCmd.Flags().String("update-type", "", "The type of update.")
	tnb_updateSolNetworkInstanceCmd.MarkFlagRequired("ns-instance-id")
	tnb_updateSolNetworkInstanceCmd.MarkFlagRequired("update-type")
	tnbCmd.AddCommand(tnb_updateSolNetworkInstanceCmd)
}
