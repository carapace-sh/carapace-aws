package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3vectors_putVectorBucketPolicyCmd = &cobra.Command{
	Use:   "put-vector-bucket-policy",
	Short: "Amazon S3 Vectors is in preview release for Amazon S3 and is subject to change.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3vectors_putVectorBucketPolicyCmd).Standalone()

	s3vectors_putVectorBucketPolicyCmd.Flags().String("policy", "", "The `JSON` that defines the policy.")
	s3vectors_putVectorBucketPolicyCmd.Flags().String("vector-bucket-arn", "", "The Amazon Resource Name (ARN) of the vector bucket.")
	s3vectors_putVectorBucketPolicyCmd.Flags().String("vector-bucket-name", "", "The name of the vector bucket.")
	s3vectors_putVectorBucketPolicyCmd.MarkFlagRequired("policy")
	s3vectorsCmd.AddCommand(s3vectors_putVectorBucketPolicyCmd)
}
