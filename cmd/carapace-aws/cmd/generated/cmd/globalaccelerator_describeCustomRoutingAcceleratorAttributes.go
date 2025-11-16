package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var globalaccelerator_describeCustomRoutingAcceleratorAttributesCmd = &cobra.Command{
	Use:   "describe-custom-routing-accelerator-attributes",
	Short: "Describe the attributes of a custom routing accelerator.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(globalaccelerator_describeCustomRoutingAcceleratorAttributesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(globalaccelerator_describeCustomRoutingAcceleratorAttributesCmd).Standalone()

		globalaccelerator_describeCustomRoutingAcceleratorAttributesCmd.Flags().String("accelerator-arn", "", "The Amazon Resource Name (ARN) of the custom routing accelerator to describe the attributes for.")
		globalaccelerator_describeCustomRoutingAcceleratorAttributesCmd.MarkFlagRequired("accelerator-arn")
	})
	globalacceleratorCmd.AddCommand(globalaccelerator_describeCustomRoutingAcceleratorAttributesCmd)
}
