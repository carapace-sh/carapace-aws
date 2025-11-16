package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconnect_removeBridgeOutputCmd = &cobra.Command{
	Use:   "remove-bridge-output",
	Short: "Removes an output from a bridge.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconnect_removeBridgeOutputCmd).Standalone()

	mediaconnect_removeBridgeOutputCmd.Flags().String("bridge-arn", "", "The Amazon Resource Name (ARN) of the bridge that you want to update.")
	mediaconnect_removeBridgeOutputCmd.Flags().String("output-name", "", "The name of the bridge output that you want to remove.")
	mediaconnect_removeBridgeOutputCmd.MarkFlagRequired("bridge-arn")
	mediaconnect_removeBridgeOutputCmd.MarkFlagRequired("output-name")
	mediaconnectCmd.AddCommand(mediaconnect_removeBridgeOutputCmd)
}
