package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_putBucketTaggingCmd = &cobra.Command{
	Use:   "put-bucket-tagging",
	Short: "This action puts tags on an Amazon S3 on Outposts bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_putBucketTaggingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3control_putBucketTaggingCmd).Standalone()

		s3control_putBucketTaggingCmd.Flags().String("account-id", "", "The Amazon Web Services account ID of the Outposts bucket.")
		s3control_putBucketTaggingCmd.Flags().String("bucket", "", "The Amazon Resource Name (ARN) of the bucket.")
		s3control_putBucketTaggingCmd.Flags().String("tagging", "", "")
		s3control_putBucketTaggingCmd.MarkFlagRequired("account-id")
		s3control_putBucketTaggingCmd.MarkFlagRequired("bucket")
		s3control_putBucketTaggingCmd.MarkFlagRequired("tagging")
	})
	s3controlCmd.AddCommand(s3control_putBucketTaggingCmd)
}
