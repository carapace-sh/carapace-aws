package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesis_enableEnhancedMonitoringCmd = &cobra.Command{
	Use:   "enable-enhanced-monitoring",
	Short: "Enables enhanced Kinesis data stream monitoring for shard-level metrics.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesis_enableEnhancedMonitoringCmd).Standalone()

	kinesis_enableEnhancedMonitoringCmd.Flags().String("shard-level-metrics", "", "List of shard-level metrics to enable.")
	kinesis_enableEnhancedMonitoringCmd.Flags().String("stream-arn", "", "The ARN of the stream.")
	kinesis_enableEnhancedMonitoringCmd.Flags().String("stream-name", "", "The name of the stream for which to enable enhanced monitoring.")
	kinesis_enableEnhancedMonitoringCmd.MarkFlagRequired("shard-level-metrics")
	kinesisCmd.AddCommand(kinesis_enableEnhancedMonitoringCmd)
}
