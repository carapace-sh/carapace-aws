package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_createBucketCmd = &cobra.Command{
	Use:   "create-bucket",
	Short: "End of support notice: As of October 1, 2025, Amazon S3 has discontinued support for Email Grantee Access Control Lists (ACLs).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_createBucketCmd).Standalone()

	s3_createBucketCmd.Flags().String("acl", "", "The canned ACL to apply to the bucket.")
	s3_createBucketCmd.Flags().String("bucket", "", "The name of the bucket to create.")
	s3_createBucketCmd.Flags().String("create-bucket-configuration", "", "The configuration information for the bucket.")
	s3_createBucketCmd.Flags().String("grant-full-control", "", "Allows grantee the read, write, read ACP, and write ACP permissions on the bucket.")
	s3_createBucketCmd.Flags().String("grant-read", "", "Allows grantee to list the objects in the bucket.")
	s3_createBucketCmd.Flags().String("grant-read-acp", "", "Allows grantee to read the bucket ACL.")
	s3_createBucketCmd.Flags().String("grant-write", "", "Allows grantee to create new objects in the bucket.")
	s3_createBucketCmd.Flags().String("grant-write-acp", "", "Allows grantee to write the ACL for the applicable bucket.")
	s3_createBucketCmd.Flags().String("object-lock-enabled-for-bucket", "", "Specifies whether you want S3 Object Lock to be enabled for the new bucket.")
	s3_createBucketCmd.Flags().String("object-ownership", "", "")
	s3_createBucketCmd.MarkFlagRequired("bucket")
	s3Cmd.AddCommand(s3_createBucketCmd)
}
