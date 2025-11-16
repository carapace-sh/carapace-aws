package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesis_describeStreamCmd = &cobra.Command{
	Use:   "describe-stream",
	Short: "Describes the specified Kinesis data stream.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesis_describeStreamCmd).Standalone()

	kinesis_describeStreamCmd.Flags().String("exclusive-start-shard-id", "", "The shard ID of the shard to start with.")
	kinesis_describeStreamCmd.Flags().String("limit", "", "The maximum number of shards to return in a single call.")
	kinesis_describeStreamCmd.Flags().String("stream-arn", "", "The ARN of the stream.")
	kinesis_describeStreamCmd.Flags().String("stream-name", "", "The name of the stream to describe.")
	kinesisCmd.AddCommand(kinesis_describeStreamCmd)
}
