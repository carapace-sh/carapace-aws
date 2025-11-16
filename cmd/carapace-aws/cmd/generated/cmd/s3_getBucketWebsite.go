package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_getBucketWebsiteCmd = &cobra.Command{
	Use:   "get-bucket-website",
	Short: "This operation is not supported for directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_getBucketWebsiteCmd).Standalone()

	s3_getBucketWebsiteCmd.Flags().String("bucket", "", "The bucket name for which to get the website configuration.")
	s3_getBucketWebsiteCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
	s3_getBucketWebsiteCmd.MarkFlagRequired("bucket")
	s3Cmd.AddCommand(s3_getBucketWebsiteCmd)
}
