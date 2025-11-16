package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3tables_deleteTableBucketMetricsConfigurationCmd = &cobra.Command{
	Use:   "delete-table-bucket-metrics-configuration",
	Short: "Deletes the metrics configuration for a table bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3tables_deleteTableBucketMetricsConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3tables_deleteTableBucketMetricsConfigurationCmd).Standalone()

		s3tables_deleteTableBucketMetricsConfigurationCmd.Flags().String("table-bucket-arn", "", "The Amazon Resource Name (ARN) of the table bucket.")
		s3tables_deleteTableBucketMetricsConfigurationCmd.MarkFlagRequired("table-bucket-arn")
	})
	s3tablesCmd.AddCommand(s3tables_deleteTableBucketMetricsConfigurationCmd)
}
