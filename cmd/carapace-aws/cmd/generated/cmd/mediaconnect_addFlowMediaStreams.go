package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconnect_addFlowMediaStreamsCmd = &cobra.Command{
	Use:   "add-flow-media-streams",
	Short: "Adds media streams to an existing flow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconnect_addFlowMediaStreamsCmd).Standalone()

	mediaconnect_addFlowMediaStreamsCmd.Flags().String("flow-arn", "", "The Amazon Resource Name (ARN) of the flow.")
	mediaconnect_addFlowMediaStreamsCmd.Flags().String("media-streams", "", "The media streams that you want to add to the flow.")
	mediaconnect_addFlowMediaStreamsCmd.MarkFlagRequired("flow-arn")
	mediaconnect_addFlowMediaStreamsCmd.MarkFlagRequired("media-streams")
	mediaconnectCmd.AddCommand(mediaconnect_addFlowMediaStreamsCmd)
}
