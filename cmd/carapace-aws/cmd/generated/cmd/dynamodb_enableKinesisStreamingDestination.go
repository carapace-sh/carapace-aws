package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodb_enableKinesisStreamingDestinationCmd = &cobra.Command{
	Use:   "enable-kinesis-streaming-destination",
	Short: "Starts table data replication to the specified Kinesis data stream at a timestamp chosen during the enable workflow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodb_enableKinesisStreamingDestinationCmd).Standalone()

	dynamodb_enableKinesisStreamingDestinationCmd.Flags().String("enable-kinesis-streaming-configuration", "", "The source for the Kinesis streaming information that is being enabled.")
	dynamodb_enableKinesisStreamingDestinationCmd.Flags().String("stream-arn", "", "The ARN for a Kinesis data stream.")
	dynamodb_enableKinesisStreamingDestinationCmd.Flags().String("table-name", "", "The name of the DynamoDB table.")
	dynamodb_enableKinesisStreamingDestinationCmd.MarkFlagRequired("stream-arn")
	dynamodb_enableKinesisStreamingDestinationCmd.MarkFlagRequired("table-name")
	dynamodbCmd.AddCommand(dynamodb_enableKinesisStreamingDestinationCmd)
}
