package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tnb_instantiateSolNetworkInstanceCmd = &cobra.Command{
	Use:   "instantiate-sol-network-instance",
	Short: "Instantiates a network instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tnb_instantiateSolNetworkInstanceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(tnb_instantiateSolNetworkInstanceCmd).Standalone()

		tnb_instantiateSolNetworkInstanceCmd.Flags().String("additional-params-for-ns", "", "Provides values for the configurable properties.")
		tnb_instantiateSolNetworkInstanceCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		tnb_instantiateSolNetworkInstanceCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		tnb_instantiateSolNetworkInstanceCmd.Flags().String("ns-instance-id", "", "ID of the network instance.")
		tnb_instantiateSolNetworkInstanceCmd.Flags().String("tags", "", "A tag is a label that you assign to an Amazon Web Services resource.")
		tnb_instantiateSolNetworkInstanceCmd.Flag("no-dry-run").Hidden = true
		tnb_instantiateSolNetworkInstanceCmd.MarkFlagRequired("ns-instance-id")
	})
	tnbCmd.AddCommand(tnb_instantiateSolNetworkInstanceCmd)
}
