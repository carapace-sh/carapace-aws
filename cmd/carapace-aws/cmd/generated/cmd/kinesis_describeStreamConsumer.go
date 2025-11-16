package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesis_describeStreamConsumerCmd = &cobra.Command{
	Use:   "describe-stream-consumer",
	Short: "To get the description of a registered consumer, provide the ARN of the consumer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesis_describeStreamConsumerCmd).Standalone()

	kinesis_describeStreamConsumerCmd.Flags().String("consumer-arn", "", "The ARN returned by Kinesis Data Streams when you registered the consumer.")
	kinesis_describeStreamConsumerCmd.Flags().String("consumer-name", "", "The name that you gave to the consumer.")
	kinesis_describeStreamConsumerCmd.Flags().String("stream-arn", "", "The ARN of the Kinesis data stream that the consumer is registered with.")
	kinesisCmd.AddCommand(kinesis_describeStreamConsumerCmd)
}
