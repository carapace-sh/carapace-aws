package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesis_deregisterStreamConsumerCmd = &cobra.Command{
	Use:   "deregister-stream-consumer",
	Short: "To deregister a consumer, provide its ARN.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesis_deregisterStreamConsumerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesis_deregisterStreamConsumerCmd).Standalone()

		kinesis_deregisterStreamConsumerCmd.Flags().String("consumer-arn", "", "The ARN returned by Kinesis Data Streams when you registered the consumer.")
		kinesis_deregisterStreamConsumerCmd.Flags().String("consumer-name", "", "The name that you gave to the consumer.")
		kinesis_deregisterStreamConsumerCmd.Flags().String("stream-arn", "", "The ARN of the Kinesis data stream that the consumer is registered with.")
	})
	kinesisCmd.AddCommand(kinesis_deregisterStreamConsumerCmd)
}
