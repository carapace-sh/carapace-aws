package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemakerMetrics_batchGetMetricsCmd = &cobra.Command{
	Use:   "batch-get-metrics",
	Short: "Used to retrieve training metrics from SageMaker.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemakerMetrics_batchGetMetricsCmd).Standalone()

	sagemakerMetrics_batchGetMetricsCmd.Flags().String("metric-queries", "", "Queries made to retrieve training metrics from SageMaker.")
	sagemakerMetrics_batchGetMetricsCmd.MarkFlagRequired("metric-queries")
	sagemakerMetricsCmd.AddCommand(sagemakerMetrics_batchGetMetricsCmd)
}
