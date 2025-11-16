package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_createBucketCmd = &cobra.Command{
	Use:   "create-bucket",
	Short: "This action creates an Amazon S3 on Outposts bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_createBucketCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3control_createBucketCmd).Standalone()

		s3control_createBucketCmd.Flags().String("acl", "", "The canned ACL to apply to the bucket.")
		s3control_createBucketCmd.Flags().String("bucket", "", "The name of the bucket.")
		s3control_createBucketCmd.Flags().String("create-bucket-configuration", "", "The configuration information for the bucket.")
		s3control_createBucketCmd.Flags().String("grant-full-control", "", "Allows grantee the read, write, read ACP, and write ACP permissions on the bucket.")
		s3control_createBucketCmd.Flags().String("grant-read", "", "Allows grantee to list the objects in the bucket.")
		s3control_createBucketCmd.Flags().String("grant-read-acp", "", "Allows grantee to read the bucket ACL.")
		s3control_createBucketCmd.Flags().String("grant-write", "", "Allows grantee to create, overwrite, and delete any object in the bucket.")
		s3control_createBucketCmd.Flags().String("grant-write-acp", "", "Allows grantee to write the ACL for the applicable bucket.")
		s3control_createBucketCmd.Flags().String("object-lock-enabled-for-bucket", "", "Specifies whether you want S3 Object Lock to be enabled for the new bucket.")
		s3control_createBucketCmd.Flags().String("outpost-id", "", "The ID of the Outposts where the bucket is being created.")
		s3control_createBucketCmd.MarkFlagRequired("bucket")
	})
	s3controlCmd.AddCommand(s3control_createBucketCmd)
}
