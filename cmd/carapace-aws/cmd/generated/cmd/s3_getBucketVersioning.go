package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_getBucketVersioningCmd = &cobra.Command{
	Use:   "get-bucket-versioning",
	Short: "This operation is not supported for directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_getBucketVersioningCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3_getBucketVersioningCmd).Standalone()

		s3_getBucketVersioningCmd.Flags().String("bucket", "", "The name of the bucket for which to get the versioning information.")
		s3_getBucketVersioningCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
		s3_getBucketVersioningCmd.MarkFlagRequired("bucket")
	})
	s3Cmd.AddCommand(s3_getBucketVersioningCmd)
}
