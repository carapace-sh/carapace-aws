package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_getBucketAclCmd = &cobra.Command{
	Use:   "get-bucket-acl",
	Short: "End of support notice: Beginning November 21, 2025, Amazon S3 will stop returning `DisplayName`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_getBucketAclCmd).Standalone()

	s3_getBucketAclCmd.Flags().String("bucket", "", "Specifies the S3 bucket whose ACL is being requested.")
	s3_getBucketAclCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
	s3_getBucketAclCmd.MarkFlagRequired("bucket")
	s3Cmd.AddCommand(s3_getBucketAclCmd)
}
