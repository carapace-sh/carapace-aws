package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_putBucketMetricsConfigurationCmd = &cobra.Command{
	Use:   "put-bucket-metrics-configuration",
	Short: "This operation is not supported for directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_putBucketMetricsConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3_putBucketMetricsConfigurationCmd).Standalone()

		s3_putBucketMetricsConfigurationCmd.Flags().String("bucket", "", "The name of the bucket for which the metrics configuration is set.")
		s3_putBucketMetricsConfigurationCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
		s3_putBucketMetricsConfigurationCmd.Flags().String("id", "", "The ID used to identify the metrics configuration.")
		s3_putBucketMetricsConfigurationCmd.Flags().String("metrics-configuration", "", "Specifies the metrics configuration.")
		s3_putBucketMetricsConfigurationCmd.MarkFlagRequired("bucket")
		s3_putBucketMetricsConfigurationCmd.MarkFlagRequired("id")
		s3_putBucketMetricsConfigurationCmd.MarkFlagRequired("metrics-configuration")
	})
	s3Cmd.AddCommand(s3_putBucketMetricsConfigurationCmd)
}
