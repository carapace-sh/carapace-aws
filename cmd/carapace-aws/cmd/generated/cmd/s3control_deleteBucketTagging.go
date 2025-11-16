package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_deleteBucketTaggingCmd = &cobra.Command{
	Use:   "delete-bucket-tagging",
	Short: "This action deletes an Amazon S3 on Outposts bucket's tags.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_deleteBucketTaggingCmd).Standalone()

	s3control_deleteBucketTaggingCmd.Flags().String("account-id", "", "The Amazon Web Services account ID of the Outposts bucket tag set to be removed.")
	s3control_deleteBucketTaggingCmd.Flags().String("bucket", "", "The bucket ARN that has the tag set to be removed.")
	s3control_deleteBucketTaggingCmd.MarkFlagRequired("account-id")
	s3control_deleteBucketTaggingCmd.MarkFlagRequired("bucket")
	s3controlCmd.AddCommand(s3control_deleteBucketTaggingCmd)
}
