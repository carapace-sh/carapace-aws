package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisvideo_describeEdgeConfigurationCmd = &cobra.Command{
	Use:   "describe-edge-configuration",
	Short: "Describes a stream’s edge configuration that was set using the `StartEdgeConfigurationUpdate` API and the latest status of the edge agent's recorder and uploader jobs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisvideo_describeEdgeConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesisvideo_describeEdgeConfigurationCmd).Standalone()

		kinesisvideo_describeEdgeConfigurationCmd.Flags().String("stream-arn", "", "The Amazon Resource Name (ARN) of the stream.")
		kinesisvideo_describeEdgeConfigurationCmd.Flags().String("stream-name", "", "The name of the stream whose edge configuration you want to update.")
	})
	kinesisvideoCmd.AddCommand(kinesisvideo_describeEdgeConfigurationCmd)
}
