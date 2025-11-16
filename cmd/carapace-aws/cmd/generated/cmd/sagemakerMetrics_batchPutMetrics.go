package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemakerMetrics_batchPutMetricsCmd = &cobra.Command{
	Use:   "batch-put-metrics",
	Short: "Used to ingest training metrics into SageMaker.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemakerMetrics_batchPutMetricsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemakerMetrics_batchPutMetricsCmd).Standalone()

		sagemakerMetrics_batchPutMetricsCmd.Flags().String("metric-data", "", "A list of raw metric values to put.")
		sagemakerMetrics_batchPutMetricsCmd.Flags().String("trial-component-name", "", "The name of the Trial Component to associate with the metrics.")
		sagemakerMetrics_batchPutMetricsCmd.MarkFlagRequired("metric-data")
		sagemakerMetrics_batchPutMetricsCmd.MarkFlagRequired("trial-component-name")
	})
	sagemakerMetricsCmd.AddCommand(sagemakerMetrics_batchPutMetricsCmd)
}
