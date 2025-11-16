package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tnb_createSolNetworkInstanceCmd = &cobra.Command{
	Use:   "create-sol-network-instance",
	Short: "Creates a network instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tnb_createSolNetworkInstanceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(tnb_createSolNetworkInstanceCmd).Standalone()

		tnb_createSolNetworkInstanceCmd.Flags().String("ns-description", "", "Network instance description.")
		tnb_createSolNetworkInstanceCmd.Flags().String("ns-name", "", "Network instance name.")
		tnb_createSolNetworkInstanceCmd.Flags().String("nsd-info-id", "", "ID for network service descriptor.")
		tnb_createSolNetworkInstanceCmd.Flags().String("tags", "", "A tag is a label that you assign to an Amazon Web Services resource.")
		tnb_createSolNetworkInstanceCmd.MarkFlagRequired("ns-name")
		tnb_createSolNetworkInstanceCmd.MarkFlagRequired("nsd-info-id")
	})
	tnbCmd.AddCommand(tnb_createSolNetworkInstanceCmd)
}
