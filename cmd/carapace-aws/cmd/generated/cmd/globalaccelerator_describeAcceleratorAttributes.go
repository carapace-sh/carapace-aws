package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var globalaccelerator_describeAcceleratorAttributesCmd = &cobra.Command{
	Use:   "describe-accelerator-attributes",
	Short: "Describe the attributes of an accelerator.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(globalaccelerator_describeAcceleratorAttributesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(globalaccelerator_describeAcceleratorAttributesCmd).Standalone()

		globalaccelerator_describeAcceleratorAttributesCmd.Flags().String("accelerator-arn", "", "The Amazon Resource Name (ARN) of the accelerator with the attributes that you want to describe.")
		globalaccelerator_describeAcceleratorAttributesCmd.MarkFlagRequired("accelerator-arn")
	})
	globalacceleratorCmd.AddCommand(globalaccelerator_describeAcceleratorAttributesCmd)
}
