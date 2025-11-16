package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesis_getRecordsCmd = &cobra.Command{
	Use:   "get-records",
	Short: "Gets data records from a Kinesis data stream's shard.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesis_getRecordsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesis_getRecordsCmd).Standalone()

		kinesis_getRecordsCmd.Flags().String("limit", "", "The maximum number of records to return.")
		kinesis_getRecordsCmd.Flags().String("shard-iterator", "", "The position in the shard from which you want to start sequentially reading data records.")
		kinesis_getRecordsCmd.Flags().String("stream-arn", "", "The ARN of the stream.")
		kinesis_getRecordsCmd.MarkFlagRequired("shard-iterator")
	})
	kinesisCmd.AddCommand(kinesis_getRecordsCmd)
}
