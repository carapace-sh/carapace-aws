package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodb_disableKinesisStreamingDestinationCmd = &cobra.Command{
	Use:   "disable-kinesis-streaming-destination",
	Short: "Stops replication from the DynamoDB table to the Kinesis data stream.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodb_disableKinesisStreamingDestinationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dynamodb_disableKinesisStreamingDestinationCmd).Standalone()

		dynamodb_disableKinesisStreamingDestinationCmd.Flags().String("enable-kinesis-streaming-configuration", "", "The source for the Kinesis streaming information that is being enabled.")
		dynamodb_disableKinesisStreamingDestinationCmd.Flags().String("stream-arn", "", "The ARN for a Kinesis data stream.")
		dynamodb_disableKinesisStreamingDestinationCmd.Flags().String("table-name", "", "The name of the DynamoDB table.")
		dynamodb_disableKinesisStreamingDestinationCmd.MarkFlagRequired("stream-arn")
		dynamodb_disableKinesisStreamingDestinationCmd.MarkFlagRequired("table-name")
	})
	dynamodbCmd.AddCommand(dynamodb_disableKinesisStreamingDestinationCmd)
}
