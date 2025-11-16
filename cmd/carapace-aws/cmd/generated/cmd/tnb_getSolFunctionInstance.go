package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tnb_getSolFunctionInstanceCmd = &cobra.Command{
	Use:   "get-sol-function-instance",
	Short: "Gets the details of a network function instance, including the instantiation state and metadata from the function package descriptor in the network function package.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tnb_getSolFunctionInstanceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(tnb_getSolFunctionInstanceCmd).Standalone()

		tnb_getSolFunctionInstanceCmd.Flags().String("vnf-instance-id", "", "ID of the network function.")
		tnb_getSolFunctionInstanceCmd.MarkFlagRequired("vnf-instance-id")
	})
	tnbCmd.AddCommand(tnb_getSolFunctionInstanceCmd)
}
