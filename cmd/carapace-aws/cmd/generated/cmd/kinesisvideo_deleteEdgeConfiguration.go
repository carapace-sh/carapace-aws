package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisvideo_deleteEdgeConfigurationCmd = &cobra.Command{
	Use:   "delete-edge-configuration",
	Short: "An asynchronous API that deletes a stream’s existing edge configuration, as well as the corresponding media from the Edge Agent.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisvideo_deleteEdgeConfigurationCmd).Standalone()

	kinesisvideo_deleteEdgeConfigurationCmd.Flags().String("stream-arn", "", "The Amazon Resource Name (ARN) of the stream.")
	kinesisvideo_deleteEdgeConfigurationCmd.Flags().String("stream-name", "", "The name of the stream from which to delete the edge configuration.")
	kinesisvideoCmd.AddCommand(kinesisvideo_deleteEdgeConfigurationCmd)
}
