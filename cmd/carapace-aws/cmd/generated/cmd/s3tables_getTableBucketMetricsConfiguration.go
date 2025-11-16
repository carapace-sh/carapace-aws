package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3tables_getTableBucketMetricsConfigurationCmd = &cobra.Command{
	Use:   "get-table-bucket-metrics-configuration",
	Short: "Gets the metrics configuration for a table bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3tables_getTableBucketMetricsConfigurationCmd).Standalone()

	s3tables_getTableBucketMetricsConfigurationCmd.Flags().String("table-bucket-arn", "", "The Amazon Resource Name (ARN) of the table bucket.")
	s3tables_getTableBucketMetricsConfigurationCmd.MarkFlagRequired("table-bucket-arn")
	s3tablesCmd.AddCommand(s3tables_getTableBucketMetricsConfigurationCmd)
}
