package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_putBucketTaggingCmd = &cobra.Command{
	Use:   "put-bucket-tagging",
	Short: "This operation is not supported for directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_putBucketTaggingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3_putBucketTaggingCmd).Standalone()

		s3_putBucketTaggingCmd.Flags().String("bucket", "", "The bucket name.")
		s3_putBucketTaggingCmd.Flags().String("checksum-algorithm", "", "Indicates the algorithm used to create the checksum for the request when you use the SDK.")
		s3_putBucketTaggingCmd.Flags().String("content-md5", "", "The Base64 encoded 128-bit `MD5` digest of the data.")
		s3_putBucketTaggingCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
		s3_putBucketTaggingCmd.Flags().String("tagging", "", "Container for the `TagSet` and `Tag` elements.")
		s3_putBucketTaggingCmd.MarkFlagRequired("bucket")
		s3_putBucketTaggingCmd.MarkFlagRequired("tagging")
	})
	s3Cmd.AddCommand(s3_putBucketTaggingCmd)
}
