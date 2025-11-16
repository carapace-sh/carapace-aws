package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_getBucketAnalyticsConfigurationCmd = &cobra.Command{
	Use:   "get-bucket-analytics-configuration",
	Short: "This operation is not supported for directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_getBucketAnalyticsConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3_getBucketAnalyticsConfigurationCmd).Standalone()

		s3_getBucketAnalyticsConfigurationCmd.Flags().String("bucket", "", "The name of the bucket from which an analytics configuration is retrieved.")
		s3_getBucketAnalyticsConfigurationCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
		s3_getBucketAnalyticsConfigurationCmd.Flags().String("id", "", "The ID that identifies the analytics configuration.")
		s3_getBucketAnalyticsConfigurationCmd.MarkFlagRequired("bucket")
		s3_getBucketAnalyticsConfigurationCmd.MarkFlagRequired("id")
	})
	s3Cmd.AddCommand(s3_getBucketAnalyticsConfigurationCmd)
}
