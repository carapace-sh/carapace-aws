package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_getBucketLifecycleCmd = &cobra.Command{
	Use:   "get-bucket-lifecycle",
	Short: "For an updated version of this API, see [GetBucketLifecycleConfiguration](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketLifecycleConfiguration.html).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_getBucketLifecycleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3_getBucketLifecycleCmd).Standalone()

		s3_getBucketLifecycleCmd.Flags().String("bucket", "", "The name of the bucket for which to get the lifecycle information.")
		s3_getBucketLifecycleCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
		s3_getBucketLifecycleCmd.MarkFlagRequired("bucket")
	})
	s3Cmd.AddCommand(s3_getBucketLifecycleCmd)
}
