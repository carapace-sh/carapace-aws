package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_getBucketPolicyCmd = &cobra.Command{
	Use:   "get-bucket-policy",
	Short: "This action gets a bucket policy for an Amazon S3 on Outposts bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_getBucketPolicyCmd).Standalone()

	s3control_getBucketPolicyCmd.Flags().String("account-id", "", "The Amazon Web Services account ID of the Outposts bucket.")
	s3control_getBucketPolicyCmd.Flags().String("bucket", "", "Specifies the bucket.")
	s3control_getBucketPolicyCmd.MarkFlagRequired("account-id")
	s3control_getBucketPolicyCmd.MarkFlagRequired("bucket")
	s3controlCmd.AddCommand(s3control_getBucketPolicyCmd)
}
