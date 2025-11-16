package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_putBucketVersioningCmd = &cobra.Command{
	Use:   "put-bucket-versioning",
	Short: "This operation is not supported for directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_putBucketVersioningCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3_putBucketVersioningCmd).Standalone()

		s3_putBucketVersioningCmd.Flags().String("bucket", "", "The bucket name.")
		s3_putBucketVersioningCmd.Flags().String("checksum-algorithm", "", "Indicates the algorithm used to create the checksum for the request when you use the SDK.")
		s3_putBucketVersioningCmd.Flags().String("content-md5", "", "&gt;The Base64 encoded 128-bit `MD5` digest of the data.")
		s3_putBucketVersioningCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
		s3_putBucketVersioningCmd.Flags().String("mfa", "", "The concatenation of the authentication device's serial number, a space, and the value that is displayed on your authentication device.")
		s3_putBucketVersioningCmd.Flags().String("versioning-configuration", "", "Container for setting the versioning state.")
		s3_putBucketVersioningCmd.MarkFlagRequired("bucket")
		s3_putBucketVersioningCmd.MarkFlagRequired("versioning-configuration")
	})
	s3Cmd.AddCommand(s3_putBucketVersioningCmd)
}
