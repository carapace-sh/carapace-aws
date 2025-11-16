package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_putBucketPolicyCmd = &cobra.Command{
	Use:   "put-bucket-policy",
	Short: "Applies an Amazon S3 bucket policy to an Amazon S3 bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_putBucketPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3_putBucketPolicyCmd).Standalone()

		s3_putBucketPolicyCmd.Flags().String("bucket", "", "The name of the bucket.")
		s3_putBucketPolicyCmd.Flags().String("checksum-algorithm", "", "Indicates the algorithm used to create the checksum for the request when you use the SDK.")
		s3_putBucketPolicyCmd.Flags().String("confirm-remove-self-bucket-access", "", "Set this parameter to true to confirm that you want to remove your permissions to change this bucket policy in the future.")
		s3_putBucketPolicyCmd.Flags().String("content-md5", "", "The MD5 hash of the request body.")
		s3_putBucketPolicyCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
		s3_putBucketPolicyCmd.Flags().String("policy", "", "The bucket policy as a JSON document.")
		s3_putBucketPolicyCmd.MarkFlagRequired("bucket")
		s3_putBucketPolicyCmd.MarkFlagRequired("policy")
	})
	s3Cmd.AddCommand(s3_putBucketPolicyCmd)
}
