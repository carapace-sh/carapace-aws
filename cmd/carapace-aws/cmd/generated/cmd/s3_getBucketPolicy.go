package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_getBucketPolicyCmd = &cobra.Command{
	Use:   "get-bucket-policy",
	Short: "Returns the policy of a specified bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_getBucketPolicyCmd).Standalone()

	s3_getBucketPolicyCmd.Flags().String("bucket", "", "The bucket name to get the bucket policy for.")
	s3_getBucketPolicyCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
	s3_getBucketPolicyCmd.MarkFlagRequired("bucket")
	s3Cmd.AddCommand(s3_getBucketPolicyCmd)
}
