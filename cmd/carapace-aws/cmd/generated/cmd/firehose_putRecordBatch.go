package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var firehose_putRecordBatchCmd = &cobra.Command{
	Use:   "put-record-batch",
	Short: "Writes multiple data records into a Firehose stream in a single call, which can achieve higher throughput per producer than when writing single records.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(firehose_putRecordBatchCmd).Standalone()

	firehose_putRecordBatchCmd.Flags().String("delivery-stream-name", "", "The name of the Firehose stream.")
	firehose_putRecordBatchCmd.Flags().String("records", "", "One or more records.")
	firehose_putRecordBatchCmd.MarkFlagRequired("delivery-stream-name")
	firehose_putRecordBatchCmd.MarkFlagRequired("records")
	firehoseCmd.AddCommand(firehose_putRecordBatchCmd)
}
