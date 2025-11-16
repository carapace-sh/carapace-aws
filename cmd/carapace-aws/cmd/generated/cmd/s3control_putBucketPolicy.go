package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_putBucketPolicyCmd = &cobra.Command{
	Use:   "put-bucket-policy",
	Short: "This action puts a bucket policy to an Amazon S3 on Outposts bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_putBucketPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3control_putBucketPolicyCmd).Standalone()

		s3control_putBucketPolicyCmd.Flags().String("account-id", "", "The Amazon Web Services account ID of the Outposts bucket.")
		s3control_putBucketPolicyCmd.Flags().String("bucket", "", "Specifies the bucket.")
		s3control_putBucketPolicyCmd.Flags().String("confirm-remove-self-bucket-access", "", "Set this parameter to true to confirm that you want to remove your permissions to change this bucket policy in the future.")
		s3control_putBucketPolicyCmd.Flags().String("policy", "", "The bucket policy as a JSON document.")
		s3control_putBucketPolicyCmd.MarkFlagRequired("account-id")
		s3control_putBucketPolicyCmd.MarkFlagRequired("bucket")
		s3control_putBucketPolicyCmd.MarkFlagRequired("policy")
	})
	s3controlCmd.AddCommand(s3control_putBucketPolicyCmd)
}
