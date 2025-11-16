package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_deleteBucketTaggingCmd = &cobra.Command{
	Use:   "delete-bucket-tagging",
	Short: "This operation is not supported for directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_deleteBucketTaggingCmd).Standalone()

	s3_deleteBucketTaggingCmd.Flags().String("bucket", "", "The bucket that has the tag set to be removed.")
	s3_deleteBucketTaggingCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
	s3_deleteBucketTaggingCmd.MarkFlagRequired("bucket")
	s3Cmd.AddCommand(s3_deleteBucketTaggingCmd)
}
