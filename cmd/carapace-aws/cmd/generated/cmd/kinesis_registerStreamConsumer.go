package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesis_registerStreamConsumerCmd = &cobra.Command{
	Use:   "register-stream-consumer",
	Short: "Registers a consumer with a Kinesis data stream.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesis_registerStreamConsumerCmd).Standalone()

	kinesis_registerStreamConsumerCmd.Flags().String("consumer-name", "", "For a given Kinesis data stream, each consumer must have a unique name.")
	kinesis_registerStreamConsumerCmd.Flags().String("stream-arn", "", "The ARN of the Kinesis data stream that you want to register the consumer with.")
	kinesis_registerStreamConsumerCmd.Flags().String("tags", "", "A set of up to 50 key-value pairs.")
	kinesis_registerStreamConsumerCmd.MarkFlagRequired("consumer-name")
	kinesis_registerStreamConsumerCmd.MarkFlagRequired("stream-arn")
	kinesisCmd.AddCommand(kinesis_registerStreamConsumerCmd)
}
