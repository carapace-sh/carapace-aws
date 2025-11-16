package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3vectors_deleteVectorBucketPolicyCmd = &cobra.Command{
	Use:   "delete-vector-bucket-policy",
	Short: "Amazon S3 Vectors is in preview release for Amazon S3 and is subject to change.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3vectors_deleteVectorBucketPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3vectors_deleteVectorBucketPolicyCmd).Standalone()

		s3vectors_deleteVectorBucketPolicyCmd.Flags().String("vector-bucket-arn", "", "The ARN of the vector bucket to delete the policy from.")
		s3vectors_deleteVectorBucketPolicyCmd.Flags().String("vector-bucket-name", "", "The name of the vector bucket to delete the policy from.")
	})
	s3vectorsCmd.AddCommand(s3vectors_deleteVectorBucketPolicyCmd)
}
