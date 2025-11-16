package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesis_decreaseStreamRetentionPeriodCmd = &cobra.Command{
	Use:   "decrease-stream-retention-period",
	Short: "Decreases the Kinesis data stream's retention period, which is the length of time data records are accessible after they are added to the stream.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesis_decreaseStreamRetentionPeriodCmd).Standalone()

	kinesis_decreaseStreamRetentionPeriodCmd.Flags().String("retention-period-hours", "", "The new retention period of the stream, in hours.")
	kinesis_decreaseStreamRetentionPeriodCmd.Flags().String("stream-arn", "", "The ARN of the stream.")
	kinesis_decreaseStreamRetentionPeriodCmd.Flags().String("stream-name", "", "The name of the stream to modify.")
	kinesis_decreaseStreamRetentionPeriodCmd.MarkFlagRequired("retention-period-hours")
	kinesisCmd.AddCommand(kinesis_decreaseStreamRetentionPeriodCmd)
}
