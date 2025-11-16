package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_putBucketWebsiteCmd = &cobra.Command{
	Use:   "put-bucket-website",
	Short: "This operation is not supported for directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_putBucketWebsiteCmd).Standalone()

	s3_putBucketWebsiteCmd.Flags().String("bucket", "", "The bucket name.")
	s3_putBucketWebsiteCmd.Flags().String("checksum-algorithm", "", "Indicates the algorithm used to create the checksum for the request when you use the SDK.")
	s3_putBucketWebsiteCmd.Flags().String("content-md5", "", "The Base64 encoded 128-bit `MD5` digest of the data.")
	s3_putBucketWebsiteCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
	s3_putBucketWebsiteCmd.Flags().String("website-configuration", "", "Container for the request.")
	s3_putBucketWebsiteCmd.MarkFlagRequired("bucket")
	s3_putBucketWebsiteCmd.MarkFlagRequired("website-configuration")
	s3Cmd.AddCommand(s3_putBucketWebsiteCmd)
}
