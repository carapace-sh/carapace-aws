package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_deleteBucketAnalyticsConfigurationCmd = &cobra.Command{
	Use:   "delete-bucket-analytics-configuration",
	Short: "This operation is not supported for directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_deleteBucketAnalyticsConfigurationCmd).Standalone()

	s3_deleteBucketAnalyticsConfigurationCmd.Flags().String("bucket", "", "The name of the bucket from which an analytics configuration is deleted.")
	s3_deleteBucketAnalyticsConfigurationCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
	s3_deleteBucketAnalyticsConfigurationCmd.Flags().String("id", "", "The ID that identifies the analytics configuration.")
	s3_deleteBucketAnalyticsConfigurationCmd.MarkFlagRequired("bucket")
	s3_deleteBucketAnalyticsConfigurationCmd.MarkFlagRequired("id")
	s3Cmd.AddCommand(s3_deleteBucketAnalyticsConfigurationCmd)
}
