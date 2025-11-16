package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodb_updateKinesisStreamingDestinationCmd = &cobra.Command{
	Use:   "update-kinesis-streaming-destination",
	Short: "The command to update the Kinesis stream destination.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodb_updateKinesisStreamingDestinationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dynamodb_updateKinesisStreamingDestinationCmd).Standalone()

		dynamodb_updateKinesisStreamingDestinationCmd.Flags().String("stream-arn", "", "The Amazon Resource Name (ARN) for the Kinesis stream input.")
		dynamodb_updateKinesisStreamingDestinationCmd.Flags().String("table-name", "", "The table name for the Kinesis streaming destination input.")
		dynamodb_updateKinesisStreamingDestinationCmd.Flags().String("update-kinesis-streaming-configuration", "", "The command to update the Kinesis stream configuration.")
		dynamodb_updateKinesisStreamingDestinationCmd.MarkFlagRequired("stream-arn")
		dynamodb_updateKinesisStreamingDestinationCmd.MarkFlagRequired("table-name")
	})
	dynamodbCmd.AddCommand(dynamodb_updateKinesisStreamingDestinationCmd)
}
