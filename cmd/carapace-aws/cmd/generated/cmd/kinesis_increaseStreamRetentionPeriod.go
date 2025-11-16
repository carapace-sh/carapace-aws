package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesis_increaseStreamRetentionPeriodCmd = &cobra.Command{
	Use:   "increase-stream-retention-period",
	Short: "Increases the Kinesis data stream's retention period, which is the length of time data records are accessible after they are added to the stream.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesis_increaseStreamRetentionPeriodCmd).Standalone()

	kinesis_increaseStreamRetentionPeriodCmd.Flags().String("retention-period-hours", "", "The new retention period of the stream, in hours.")
	kinesis_increaseStreamRetentionPeriodCmd.Flags().String("stream-arn", "", "The ARN of the stream.")
	kinesis_increaseStreamRetentionPeriodCmd.Flags().String("stream-name", "", "The name of the stream to modify.")
	kinesis_increaseStreamRetentionPeriodCmd.MarkFlagRequired("retention-period-hours")
	kinesisCmd.AddCommand(kinesis_increaseStreamRetentionPeriodCmd)
}
