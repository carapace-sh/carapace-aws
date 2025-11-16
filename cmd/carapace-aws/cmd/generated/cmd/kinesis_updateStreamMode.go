package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesis_updateStreamModeCmd = &cobra.Command{
	Use:   "update-stream-mode",
	Short: "Updates the capacity mode of the data stream.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesis_updateStreamModeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesis_updateStreamModeCmd).Standalone()

		kinesis_updateStreamModeCmd.Flags().String("stream-arn", "", "Specifies the ARN of the data stream whose capacity mode you want to update.")
		kinesis_updateStreamModeCmd.Flags().String("stream-mode-details", "", "Specifies the capacity mode to which you want to set your data stream.")
		kinesis_updateStreamModeCmd.Flags().String("warm-throughput-mi-bps", "", "The target warm throughput in MB/s that the stream should be scaled to handle.")
		kinesis_updateStreamModeCmd.MarkFlagRequired("stream-arn")
		kinesis_updateStreamModeCmd.MarkFlagRequired("stream-mode-details")
	})
	kinesisCmd.AddCommand(kinesis_updateStreamModeCmd)
}
