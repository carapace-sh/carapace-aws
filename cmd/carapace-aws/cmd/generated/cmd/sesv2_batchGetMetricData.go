package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_batchGetMetricDataCmd = &cobra.Command{
	Use:   "batch-get-metric-data",
	Short: "Retrieves batches of metric data collected based on your sending activity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_batchGetMetricDataCmd).Standalone()

	sesv2_batchGetMetricDataCmd.Flags().String("queries", "", "A list of queries for metrics to be retrieved.")
	sesv2_batchGetMetricDataCmd.MarkFlagRequired("queries")
	sesv2Cmd.AddCommand(sesv2_batchGetMetricDataCmd)
}
