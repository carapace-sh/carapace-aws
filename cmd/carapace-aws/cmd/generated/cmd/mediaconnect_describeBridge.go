package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconnect_describeBridgeCmd = &cobra.Command{
	Use:   "describe-bridge",
	Short: "Displays the details of a bridge.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconnect_describeBridgeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediaconnect_describeBridgeCmd).Standalone()

		mediaconnect_describeBridgeCmd.Flags().String("bridge-arn", "", "The Amazon Resource Name (ARN) of the bridge that you want to describe.")
		mediaconnect_describeBridgeCmd.MarkFlagRequired("bridge-arn")
	})
	mediaconnectCmd.AddCommand(mediaconnect_describeBridgeCmd)
}
