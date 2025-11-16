package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_deleteBucketWebsiteCmd = &cobra.Command{
	Use:   "delete-bucket-website",
	Short: "This operation is not supported for directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_deleteBucketWebsiteCmd).Standalone()

	s3_deleteBucketWebsiteCmd.Flags().String("bucket", "", "The bucket name for which you want to remove the website configuration.")
	s3_deleteBucketWebsiteCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
	s3_deleteBucketWebsiteCmd.MarkFlagRequired("bucket")
	s3Cmd.AddCommand(s3_deleteBucketWebsiteCmd)
}
