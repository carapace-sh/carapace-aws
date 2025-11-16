package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesis_updateStreamWarmThroughputCmd = &cobra.Command{
	Use:   "update-stream-warm-throughput",
	Short: "Updates the warm throughput configuration for the specified Amazon Kinesis Data Streams on-demand data stream.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesis_updateStreamWarmThroughputCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesis_updateStreamWarmThroughputCmd).Standalone()

		kinesis_updateStreamWarmThroughputCmd.Flags().String("stream-arn", "", "The ARN of the stream to be updated.")
		kinesis_updateStreamWarmThroughputCmd.Flags().String("stream-name", "", "The name of the stream to be updated.")
		kinesis_updateStreamWarmThroughputCmd.Flags().String("warm-throughput-mi-bps", "", "The target warm throughput in MB/s that the stream should be scaled to handle.")
		kinesis_updateStreamWarmThroughputCmd.MarkFlagRequired("warm-throughput-mi-bps")
	})
	kinesisCmd.AddCommand(kinesis_updateStreamWarmThroughputCmd)
}
