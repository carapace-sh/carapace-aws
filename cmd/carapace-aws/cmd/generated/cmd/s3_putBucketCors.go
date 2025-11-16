package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_putBucketCorsCmd = &cobra.Command{
	Use:   "put-bucket-cors",
	Short: "This operation is not supported for directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_putBucketCorsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3_putBucketCorsCmd).Standalone()

		s3_putBucketCorsCmd.Flags().String("bucket", "", "Specifies the bucket impacted by the `cors`configuration.")
		s3_putBucketCorsCmd.Flags().String("checksum-algorithm", "", "Indicates the algorithm used to create the checksum for the request when you use the SDK.")
		s3_putBucketCorsCmd.Flags().String("content-md5", "", "The Base64 encoded 128-bit `MD5` digest of the data.")
		s3_putBucketCorsCmd.Flags().String("corsconfiguration", "", "Describes the cross-origin access configuration for objects in an Amazon S3 bucket.")
		s3_putBucketCorsCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
		s3_putBucketCorsCmd.MarkFlagRequired("bucket")
		s3_putBucketCorsCmd.MarkFlagRequired("corsconfiguration")
	})
	s3Cmd.AddCommand(s3_putBucketCorsCmd)
}
