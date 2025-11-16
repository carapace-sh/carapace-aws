package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var firehose_putRecordCmd = &cobra.Command{
	Use:   "put-record",
	Short: "Writes a single data record into an Firehose stream.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(firehose_putRecordCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(firehose_putRecordCmd).Standalone()

		firehose_putRecordCmd.Flags().String("delivery-stream-name", "", "The name of the Firehose stream.")
		firehose_putRecordCmd.Flags().String("record", "", "The record.")
		firehose_putRecordCmd.MarkFlagRequired("delivery-stream-name")
		firehose_putRecordCmd.MarkFlagRequired("record")
	})
	firehoseCmd.AddCommand(firehose_putRecordCmd)
}
