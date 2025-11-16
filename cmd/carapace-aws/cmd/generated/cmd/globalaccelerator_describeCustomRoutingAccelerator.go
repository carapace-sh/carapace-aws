package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var globalaccelerator_describeCustomRoutingAcceleratorCmd = &cobra.Command{
	Use:   "describe-custom-routing-accelerator",
	Short: "Describe a custom routing accelerator.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(globalaccelerator_describeCustomRoutingAcceleratorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(globalaccelerator_describeCustomRoutingAcceleratorCmd).Standalone()

		globalaccelerator_describeCustomRoutingAcceleratorCmd.Flags().String("accelerator-arn", "", "The Amazon Resource Name (ARN) of the accelerator to describe.")
		globalaccelerator_describeCustomRoutingAcceleratorCmd.MarkFlagRequired("accelerator-arn")
	})
	globalacceleratorCmd.AddCommand(globalaccelerator_describeCustomRoutingAcceleratorCmd)
}
