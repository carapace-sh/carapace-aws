package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisvideo_describeStreamCmd = &cobra.Command{
	Use:   "describe-stream",
	Short: "Returns the most current information about the specified stream.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisvideo_describeStreamCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesisvideo_describeStreamCmd).Standalone()

		kinesisvideo_describeStreamCmd.Flags().String("stream-arn", "", "The Amazon Resource Name (ARN) of the stream.")
		kinesisvideo_describeStreamCmd.Flags().String("stream-name", "", "The name of the stream.")
	})
	kinesisvideoCmd.AddCommand(kinesisvideo_describeStreamCmd)
}
