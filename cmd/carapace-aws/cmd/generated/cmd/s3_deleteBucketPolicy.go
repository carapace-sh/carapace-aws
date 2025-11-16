package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_deleteBucketPolicyCmd = &cobra.Command{
	Use:   "delete-bucket-policy",
	Short: "Deletes the policy of a specified bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_deleteBucketPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3_deleteBucketPolicyCmd).Standalone()

		s3_deleteBucketPolicyCmd.Flags().String("bucket", "", "The bucket name.")
		s3_deleteBucketPolicyCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
		s3_deleteBucketPolicyCmd.MarkFlagRequired("bucket")
	})
	s3Cmd.AddCommand(s3_deleteBucketPolicyCmd)
}
