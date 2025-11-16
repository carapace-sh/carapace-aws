package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesis_describeStreamSummaryCmd = &cobra.Command{
	Use:   "describe-stream-summary",
	Short: "Provides a summarized description of the specified Kinesis data stream without the shard list.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesis_describeStreamSummaryCmd).Standalone()

	kinesis_describeStreamSummaryCmd.Flags().String("stream-arn", "", "The ARN of the stream.")
	kinesis_describeStreamSummaryCmd.Flags().String("stream-name", "", "The name of the stream to describe.")
	kinesisCmd.AddCommand(kinesis_describeStreamSummaryCmd)
}
