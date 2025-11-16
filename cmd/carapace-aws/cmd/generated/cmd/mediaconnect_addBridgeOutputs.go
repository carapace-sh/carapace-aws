package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconnect_addBridgeOutputsCmd = &cobra.Command{
	Use:   "add-bridge-outputs",
	Short: "Adds outputs to an existing bridge.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconnect_addBridgeOutputsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediaconnect_addBridgeOutputsCmd).Standalone()

		mediaconnect_addBridgeOutputsCmd.Flags().String("bridge-arn", "", "The Amazon Resource Name (ARN) of the bridge that you want to update.")
		mediaconnect_addBridgeOutputsCmd.Flags().String("outputs", "", "The outputs that you want to add to this bridge.")
		mediaconnect_addBridgeOutputsCmd.MarkFlagRequired("bridge-arn")
		mediaconnect_addBridgeOutputsCmd.MarkFlagRequired("outputs")
	})
	mediaconnectCmd.AddCommand(mediaconnect_addBridgeOutputsCmd)
}
