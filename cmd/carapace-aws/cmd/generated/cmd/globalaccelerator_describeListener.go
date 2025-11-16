package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var globalaccelerator_describeListenerCmd = &cobra.Command{
	Use:   "describe-listener",
	Short: "Describe a listener.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(globalaccelerator_describeListenerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(globalaccelerator_describeListenerCmd).Standalone()

		globalaccelerator_describeListenerCmd.Flags().String("listener-arn", "", "The Amazon Resource Name (ARN) of the listener to describe.")
		globalaccelerator_describeListenerCmd.MarkFlagRequired("listener-arn")
	})
	globalacceleratorCmd.AddCommand(globalaccelerator_describeListenerCmd)
}
