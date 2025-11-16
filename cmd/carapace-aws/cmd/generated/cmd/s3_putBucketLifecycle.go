package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_putBucketLifecycleCmd = &cobra.Command{
	Use:   "put-bucket-lifecycle",
	Short: "This operation is not supported for directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_putBucketLifecycleCmd).Standalone()

	s3_putBucketLifecycleCmd.Flags().String("bucket", "", "")
	s3_putBucketLifecycleCmd.Flags().String("checksum-algorithm", "", "Indicates the algorithm used to create the checksum for the request when you use the SDK.")
	s3_putBucketLifecycleCmd.Flags().String("content-md5", "", "For requests made using the Amazon Web Services Command Line Interface (CLI) or Amazon Web Services SDKs, this field is calculated automatically.")
	s3_putBucketLifecycleCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
	s3_putBucketLifecycleCmd.Flags().String("lifecycle-configuration", "", "")
	s3_putBucketLifecycleCmd.MarkFlagRequired("bucket")
	s3Cmd.AddCommand(s3_putBucketLifecycleCmd)
}
