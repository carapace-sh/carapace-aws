package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_getBucketTaggingCmd = &cobra.Command{
	Use:   "get-bucket-tagging",
	Short: "This action gets an Amazon S3 on Outposts bucket's tags.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_getBucketTaggingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3control_getBucketTaggingCmd).Standalone()

		s3control_getBucketTaggingCmd.Flags().String("account-id", "", "The Amazon Web Services account ID of the Outposts bucket.")
		s3control_getBucketTaggingCmd.Flags().String("bucket", "", "Specifies the bucket.")
		s3control_getBucketTaggingCmd.MarkFlagRequired("account-id")
		s3control_getBucketTaggingCmd.MarkFlagRequired("bucket")
	})
	s3controlCmd.AddCommand(s3control_getBucketTaggingCmd)
}
