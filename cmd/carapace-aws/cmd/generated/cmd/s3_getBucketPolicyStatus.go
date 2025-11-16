package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_getBucketPolicyStatusCmd = &cobra.Command{
	Use:   "get-bucket-policy-status",
	Short: "This operation is not supported for directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_getBucketPolicyStatusCmd).Standalone()

	s3_getBucketPolicyStatusCmd.Flags().String("bucket", "", "The name of the Amazon S3 bucket whose policy status you want to retrieve.")
	s3_getBucketPolicyStatusCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
	s3_getBucketPolicyStatusCmd.MarkFlagRequired("bucket")
	s3Cmd.AddCommand(s3_getBucketPolicyStatusCmd)
}
