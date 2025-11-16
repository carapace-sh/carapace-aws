package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconnect_removeFlowMediaStreamCmd = &cobra.Command{
	Use:   "remove-flow-media-stream",
	Short: "Removes a media stream from a flow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconnect_removeFlowMediaStreamCmd).Standalone()

	mediaconnect_removeFlowMediaStreamCmd.Flags().String("flow-arn", "", "The Amazon Resource Name (ARN) of the flow that you want to update.")
	mediaconnect_removeFlowMediaStreamCmd.Flags().String("media-stream-name", "", "The name of the media stream that you want to remove.")
	mediaconnect_removeFlowMediaStreamCmd.MarkFlagRequired("flow-arn")
	mediaconnect_removeFlowMediaStreamCmd.MarkFlagRequired("media-stream-name")
	mediaconnectCmd.AddCommand(mediaconnect_removeFlowMediaStreamCmd)
}
