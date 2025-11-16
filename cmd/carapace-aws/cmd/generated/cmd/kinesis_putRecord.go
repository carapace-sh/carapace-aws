package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesis_putRecordCmd = &cobra.Command{
	Use:   "put-record",
	Short: "Writes a single data record into an Amazon Kinesis data stream.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesis_putRecordCmd).Standalone()

	kinesis_putRecordCmd.Flags().String("data", "", "The data blob to put into the record, which is base64-encoded when the blob is serialized.")
	kinesis_putRecordCmd.Flags().String("explicit-hash-key", "", "The hash value used to explicitly determine the shard the data record is assigned to by overriding the partition key hash.")
	kinesis_putRecordCmd.Flags().String("partition-key", "", "Determines which shard in the stream the data record is assigned to.")
	kinesis_putRecordCmd.Flags().String("sequence-number-for-ordering", "", "Guarantees strictly increasing sequence numbers, for puts from the same client and to the same partition key.")
	kinesis_putRecordCmd.Flags().String("stream-arn", "", "The ARN of the stream.")
	kinesis_putRecordCmd.Flags().String("stream-name", "", "The name of the stream to put the data record into.")
	kinesis_putRecordCmd.MarkFlagRequired("data")
	kinesis_putRecordCmd.MarkFlagRequired("partition-key")
	kinesisCmd.AddCommand(kinesis_putRecordCmd)
}
