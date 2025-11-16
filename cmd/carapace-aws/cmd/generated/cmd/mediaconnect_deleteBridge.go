package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconnect_deleteBridgeCmd = &cobra.Command{
	Use:   "delete-bridge",
	Short: "Deletes a bridge.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconnect_deleteBridgeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediaconnect_deleteBridgeCmd).Standalone()

		mediaconnect_deleteBridgeCmd.Flags().String("bridge-arn", "", "The Amazon Resource Name (ARN) of the bridge that you want to delete.")
		mediaconnect_deleteBridgeCmd.MarkFlagRequired("bridge-arn")
	})
	mediaconnectCmd.AddCommand(mediaconnect_deleteBridgeCmd)
}
