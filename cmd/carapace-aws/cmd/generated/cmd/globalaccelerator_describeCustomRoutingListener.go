package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var globalaccelerator_describeCustomRoutingListenerCmd = &cobra.Command{
	Use:   "describe-custom-routing-listener",
	Short: "The description of a listener for a custom routing accelerator.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(globalaccelerator_describeCustomRoutingListenerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(globalaccelerator_describeCustomRoutingListenerCmd).Standalone()

		globalaccelerator_describeCustomRoutingListenerCmd.Flags().String("listener-arn", "", "The Amazon Resource Name (ARN) of the listener to describe.")
		globalaccelerator_describeCustomRoutingListenerCmd.MarkFlagRequired("listener-arn")
	})
	globalacceleratorCmd.AddCommand(globalaccelerator_describeCustomRoutingListenerCmd)
}
