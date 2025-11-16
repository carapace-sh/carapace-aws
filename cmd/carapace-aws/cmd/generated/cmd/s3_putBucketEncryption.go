package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_putBucketEncryptionCmd = &cobra.Command{
	Use:   "put-bucket-encryption",
	Short: "This operation configures default encryption and Amazon S3 Bucket Keys for an existing bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_putBucketEncryptionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3_putBucketEncryptionCmd).Standalone()

		s3_putBucketEncryptionCmd.Flags().String("bucket", "", "Specifies default encryption for a bucket using server-side encryption with different key options.")
		s3_putBucketEncryptionCmd.Flags().String("checksum-algorithm", "", "Indicates the algorithm used to create the checksum for the request when you use the SDK.")
		s3_putBucketEncryptionCmd.Flags().String("content-md5", "", "The Base64 encoded 128-bit `MD5` digest of the server-side encryption configuration.")
		s3_putBucketEncryptionCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
		s3_putBucketEncryptionCmd.Flags().String("server-side-encryption-configuration", "", "")
		s3_putBucketEncryptionCmd.MarkFlagRequired("bucket")
		s3_putBucketEncryptionCmd.MarkFlagRequired("server-side-encryption-configuration")
	})
	s3Cmd.AddCommand(s3_putBucketEncryptionCmd)
}
