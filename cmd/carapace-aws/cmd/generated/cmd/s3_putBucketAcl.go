package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_putBucketAclCmd = &cobra.Command{
	Use:   "put-bucket-acl",
	Short: "End of support notice: As of October 1, 2025, Amazon S3 has discontinued support for Email Grantee Access Control Lists (ACLs).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_putBucketAclCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3_putBucketAclCmd).Standalone()

		s3_putBucketAclCmd.Flags().String("access-control-policy", "", "Contains the elements that set the ACL permissions for an object per grantee.")
		s3_putBucketAclCmd.Flags().String("acl", "", "The canned ACL to apply to the bucket.")
		s3_putBucketAclCmd.Flags().String("bucket", "", "The bucket to which to apply the ACL.")
		s3_putBucketAclCmd.Flags().String("checksum-algorithm", "", "Indicates the algorithm used to create the checksum for the request when you use the SDK.")
		s3_putBucketAclCmd.Flags().String("content-md5", "", "The Base64 encoded 128-bit `MD5` digest of the data.")
		s3_putBucketAclCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
		s3_putBucketAclCmd.Flags().String("grant-full-control", "", "Allows grantee the read, write, read ACP, and write ACP permissions on the bucket.")
		s3_putBucketAclCmd.Flags().String("grant-read", "", "Allows grantee to list the objects in the bucket.")
		s3_putBucketAclCmd.Flags().String("grant-read-acp", "", "Allows grantee to read the bucket ACL.")
		s3_putBucketAclCmd.Flags().String("grant-write", "", "Allows grantee to create new objects in the bucket.")
		s3_putBucketAclCmd.Flags().String("grant-write-acp", "", "Allows grantee to write the ACL for the applicable bucket.")
		s3_putBucketAclCmd.MarkFlagRequired("bucket")
	})
	s3Cmd.AddCommand(s3_putBucketAclCmd)
}
