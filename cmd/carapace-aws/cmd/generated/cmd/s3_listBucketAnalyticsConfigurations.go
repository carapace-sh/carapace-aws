package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_listBucketAnalyticsConfigurationsCmd = &cobra.Command{
	Use:   "list-bucket-analytics-configurations",
	Short: "This operation is not supported for directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_listBucketAnalyticsConfigurationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3_listBucketAnalyticsConfigurationsCmd).Standalone()

		s3_listBucketAnalyticsConfigurationsCmd.Flags().String("bucket", "", "The name of the bucket from which analytics configurations are retrieved.")
		s3_listBucketAnalyticsConfigurationsCmd.Flags().String("continuation-token", "", "The `ContinuationToken` that represents a placeholder from where this request should begin.")
		s3_listBucketAnalyticsConfigurationsCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
		s3_listBucketAnalyticsConfigurationsCmd.MarkFlagRequired("bucket")
	})
	s3Cmd.AddCommand(s3_listBucketAnalyticsConfigurationsCmd)
}
