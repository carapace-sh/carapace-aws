package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisvideo_updateDataRetentionCmd = &cobra.Command{
	Use:   "update-data-retention",
	Short: "Increases or decreases the stream's data retention period by the value that you specify.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisvideo_updateDataRetentionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesisvideo_updateDataRetentionCmd).Standalone()

		kinesisvideo_updateDataRetentionCmd.Flags().String("current-version", "", "The version of the stream whose retention period you want to change.")
		kinesisvideo_updateDataRetentionCmd.Flags().String("data-retention-change-in-hours", "", "The number of hours to adjust the current retention by.")
		kinesisvideo_updateDataRetentionCmd.Flags().String("operation", "", "Indicates whether you want to increase or decrease the retention period.")
		kinesisvideo_updateDataRetentionCmd.Flags().String("stream-arn", "", "The Amazon Resource Name (ARN) of the stream whose retention period you want to change.")
		kinesisvideo_updateDataRetentionCmd.Flags().String("stream-name", "", "The name of the stream whose retention period you want to change.")
		kinesisvideo_updateDataRetentionCmd.MarkFlagRequired("current-version")
		kinesisvideo_updateDataRetentionCmd.MarkFlagRequired("data-retention-change-in-hours")
		kinesisvideo_updateDataRetentionCmd.MarkFlagRequired("operation")
	})
	kinesisvideoCmd.AddCommand(kinesisvideo_updateDataRetentionCmd)
}
