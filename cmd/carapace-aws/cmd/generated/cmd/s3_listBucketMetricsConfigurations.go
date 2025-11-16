package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_listBucketMetricsConfigurationsCmd = &cobra.Command{
	Use:   "list-bucket-metrics-configurations",
	Short: "This operation is not supported for directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_listBucketMetricsConfigurationsCmd).Standalone()

	s3_listBucketMetricsConfigurationsCmd.Flags().String("bucket", "", "The name of the bucket containing the metrics configurations to retrieve.")
	s3_listBucketMetricsConfigurationsCmd.Flags().String("continuation-token", "", "The marker that is used to continue a metrics configuration listing that has been truncated.")
	s3_listBucketMetricsConfigurationsCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
	s3_listBucketMetricsConfigurationsCmd.MarkFlagRequired("bucket")
	s3Cmd.AddCommand(s3_listBucketMetricsConfigurationsCmd)
}
