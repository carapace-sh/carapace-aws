package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodb_describeKinesisStreamingDestinationCmd = &cobra.Command{
	Use:   "describe-kinesis-streaming-destination",
	Short: "Returns information about the status of Kinesis streaming.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodb_describeKinesisStreamingDestinationCmd).Standalone()

	dynamodb_describeKinesisStreamingDestinationCmd.Flags().String("table-name", "", "The name of the table being described.")
	dynamodb_describeKinesisStreamingDestinationCmd.MarkFlagRequired("table-name")
	dynamodbCmd.AddCommand(dynamodb_describeKinesisStreamingDestinationCmd)
}
