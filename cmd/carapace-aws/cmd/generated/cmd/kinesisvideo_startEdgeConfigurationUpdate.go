package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisvideo_startEdgeConfigurationUpdateCmd = &cobra.Command{
	Use:   "start-edge-configuration-update",
	Short: "An asynchronous API that updates a stream’s existing edge configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisvideo_startEdgeConfigurationUpdateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesisvideo_startEdgeConfigurationUpdateCmd).Standalone()

		kinesisvideo_startEdgeConfigurationUpdateCmd.Flags().String("edge-config", "", "The edge configuration details required to invoke the update process.")
		kinesisvideo_startEdgeConfigurationUpdateCmd.Flags().String("stream-arn", "", "The Amazon Resource Name (ARN) of the stream.")
		kinesisvideo_startEdgeConfigurationUpdateCmd.Flags().String("stream-name", "", "The name of the stream whose edge configuration you want to update.")
		kinesisvideo_startEdgeConfigurationUpdateCmd.MarkFlagRequired("edge-config")
	})
	kinesisvideoCmd.AddCommand(kinesisvideo_startEdgeConfigurationUpdateCmd)
}
