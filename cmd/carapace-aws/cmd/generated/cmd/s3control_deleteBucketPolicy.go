package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_deleteBucketPolicyCmd = &cobra.Command{
	Use:   "delete-bucket-policy",
	Short: "This action deletes an Amazon S3 on Outposts bucket policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_deleteBucketPolicyCmd).Standalone()

	s3control_deleteBucketPolicyCmd.Flags().String("account-id", "", "The account ID of the Outposts bucket.")
	s3control_deleteBucketPolicyCmd.Flags().String("bucket", "", "Specifies the bucket.")
	s3control_deleteBucketPolicyCmd.MarkFlagRequired("account-id")
	s3control_deleteBucketPolicyCmd.MarkFlagRequired("bucket")
	s3controlCmd.AddCommand(s3control_deleteBucketPolicyCmd)
}
