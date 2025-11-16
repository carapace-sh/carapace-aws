package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_putBucketAnalyticsConfigurationCmd = &cobra.Command{
	Use:   "put-bucket-analytics-configuration",
	Short: "This operation is not supported for directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_putBucketAnalyticsConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3_putBucketAnalyticsConfigurationCmd).Standalone()

		s3_putBucketAnalyticsConfigurationCmd.Flags().String("analytics-configuration", "", "The configuration and any analyses for the analytics filter.")
		s3_putBucketAnalyticsConfigurationCmd.Flags().String("bucket", "", "The name of the bucket to which an analytics configuration is stored.")
		s3_putBucketAnalyticsConfigurationCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
		s3_putBucketAnalyticsConfigurationCmd.Flags().String("id", "", "The ID that identifies the analytics configuration.")
		s3_putBucketAnalyticsConfigurationCmd.MarkFlagRequired("analytics-configuration")
		s3_putBucketAnalyticsConfigurationCmd.MarkFlagRequired("bucket")
		s3_putBucketAnalyticsConfigurationCmd.MarkFlagRequired("id")
	})
	s3Cmd.AddCommand(s3_putBucketAnalyticsConfigurationCmd)
}
