package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_getBucketTaggingCmd = &cobra.Command{
	Use:   "get-bucket-tagging",
	Short: "This operation is not supported for directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_getBucketTaggingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3_getBucketTaggingCmd).Standalone()

		s3_getBucketTaggingCmd.Flags().String("bucket", "", "The name of the bucket for which to get the tagging information.")
		s3_getBucketTaggingCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
		s3_getBucketTaggingCmd.MarkFlagRequired("bucket")
	})
	s3Cmd.AddCommand(s3_getBucketTaggingCmd)
}
