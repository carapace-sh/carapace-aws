package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesis_updateMaxRecordSizeCmd = &cobra.Command{
	Use:   "update-max-record-size",
	Short: "This allows you to update the `MaxRecordSize` of a single record that you can write to, and read from a stream.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesis_updateMaxRecordSizeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesis_updateMaxRecordSizeCmd).Standalone()

		kinesis_updateMaxRecordSizeCmd.Flags().String("max-record-size-in-ki-b", "", "The maximum record size of a single record in KiB that you can write to, and read from a stream.")
		kinesis_updateMaxRecordSizeCmd.Flags().String("stream-arn", "", "The Amazon Resource Name (ARN) of the stream for the `MaxRecordSize` update.")
		kinesis_updateMaxRecordSizeCmd.MarkFlagRequired("max-record-size-in-ki-b")
	})
	kinesisCmd.AddCommand(kinesis_updateMaxRecordSizeCmd)
}
