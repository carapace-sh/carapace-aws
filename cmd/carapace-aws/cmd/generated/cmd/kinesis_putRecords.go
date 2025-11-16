package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesis_putRecordsCmd = &cobra.Command{
	Use:   "put-records",
	Short: "Writes multiple data records into a Kinesis data stream in a single call (also referred to as a `PutRecords` request).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesis_putRecordsCmd).Standalone()

	kinesis_putRecordsCmd.Flags().String("records", "", "The records associated with the request.")
	kinesis_putRecordsCmd.Flags().String("stream-arn", "", "The ARN of the stream.")
	kinesis_putRecordsCmd.Flags().String("stream-name", "", "The stream name associated with the request.")
	kinesis_putRecordsCmd.MarkFlagRequired("records")
	kinesisCmd.AddCommand(kinesis_putRecordsCmd)
}
