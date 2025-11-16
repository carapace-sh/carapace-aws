package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconnect_removeBridgeSourceCmd = &cobra.Command{
	Use:   "remove-bridge-source",
	Short: "Removes a source from a bridge.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconnect_removeBridgeSourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediaconnect_removeBridgeSourceCmd).Standalone()

		mediaconnect_removeBridgeSourceCmd.Flags().String("bridge-arn", "", "The Amazon Resource Name (ARN) of the bridge that you want to update.")
		mediaconnect_removeBridgeSourceCmd.Flags().String("source-name", "", "The name of the bridge source that you want to remove.")
		mediaconnect_removeBridgeSourceCmd.MarkFlagRequired("bridge-arn")
		mediaconnect_removeBridgeSourceCmd.MarkFlagRequired("source-name")
	})
	mediaconnectCmd.AddCommand(mediaconnect_removeBridgeSourceCmd)
}
