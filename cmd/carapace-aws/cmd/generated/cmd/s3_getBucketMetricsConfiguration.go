package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_getBucketMetricsConfigurationCmd = &cobra.Command{
	Use:   "get-bucket-metrics-configuration",
	Short: "This operation is not supported for directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_getBucketMetricsConfigurationCmd).Standalone()

	s3_getBucketMetricsConfigurationCmd.Flags().String("bucket", "", "The name of the bucket containing the metrics configuration to retrieve.")
	s3_getBucketMetricsConfigurationCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
	s3_getBucketMetricsConfigurationCmd.Flags().String("id", "", "The ID used to identify the metrics configuration.")
	s3_getBucketMetricsConfigurationCmd.MarkFlagRequired("bucket")
	s3_getBucketMetricsConfigurationCmd.MarkFlagRequired("id")
	s3Cmd.AddCommand(s3_getBucketMetricsConfigurationCmd)
}
