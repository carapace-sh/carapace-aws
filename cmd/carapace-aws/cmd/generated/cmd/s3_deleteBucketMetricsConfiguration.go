package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_deleteBucketMetricsConfigurationCmd = &cobra.Command{
	Use:   "delete-bucket-metrics-configuration",
	Short: "This operation is not supported for directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_deleteBucketMetricsConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3_deleteBucketMetricsConfigurationCmd).Standalone()

		s3_deleteBucketMetricsConfigurationCmd.Flags().String("bucket", "", "The name of the bucket containing the metrics configuration to delete.")
		s3_deleteBucketMetricsConfigurationCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
		s3_deleteBucketMetricsConfigurationCmd.Flags().String("id", "", "The ID used to identify the metrics configuration.")
		s3_deleteBucketMetricsConfigurationCmd.MarkFlagRequired("bucket")
		s3_deleteBucketMetricsConfigurationCmd.MarkFlagRequired("id")
	})
	s3Cmd.AddCommand(s3_deleteBucketMetricsConfigurationCmd)
}
