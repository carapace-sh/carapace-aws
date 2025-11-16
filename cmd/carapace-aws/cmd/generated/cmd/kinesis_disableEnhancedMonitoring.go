package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesis_disableEnhancedMonitoringCmd = &cobra.Command{
	Use:   "disable-enhanced-monitoring",
	Short: "Disables enhanced monitoring.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesis_disableEnhancedMonitoringCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesis_disableEnhancedMonitoringCmd).Standalone()

		kinesis_disableEnhancedMonitoringCmd.Flags().String("shard-level-metrics", "", "List of shard-level metrics to disable.")
		kinesis_disableEnhancedMonitoringCmd.Flags().String("stream-arn", "", "The ARN of the stream.")
		kinesis_disableEnhancedMonitoringCmd.Flags().String("stream-name", "", "The name of the Kinesis data stream for which to disable enhanced monitoring.")
		kinesis_disableEnhancedMonitoringCmd.MarkFlagRequired("shard-level-metrics")
	})
	kinesisCmd.AddCommand(kinesis_disableEnhancedMonitoringCmd)
}
