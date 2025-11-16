package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var globalaccelerator_describeAcceleratorCmd = &cobra.Command{
	Use:   "describe-accelerator",
	Short: "Describe an accelerator.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(globalaccelerator_describeAcceleratorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(globalaccelerator_describeAcceleratorCmd).Standalone()

		globalaccelerator_describeAcceleratorCmd.Flags().String("accelerator-arn", "", "The Amazon Resource Name (ARN) of the accelerator to describe.")
		globalaccelerator_describeAcceleratorCmd.MarkFlagRequired("accelerator-arn")
	})
	globalacceleratorCmd.AddCommand(globalaccelerator_describeAcceleratorCmd)
}
