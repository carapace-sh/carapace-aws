package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_getBucketVersioningCmd = &cobra.Command{
	Use:   "get-bucket-versioning",
	Short: "This operation returns the versioning state for S3 on Outposts buckets only.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_getBucketVersioningCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3control_getBucketVersioningCmd).Standalone()

		s3control_getBucketVersioningCmd.Flags().String("account-id", "", "The Amazon Web Services account ID of the S3 on Outposts bucket.")
		s3control_getBucketVersioningCmd.Flags().String("bucket", "", "The S3 on Outposts bucket to return the versioning state for.")
		s3control_getBucketVersioningCmd.MarkFlagRequired("account-id")
		s3control_getBucketVersioningCmd.MarkFlagRequired("bucket")
	})
	s3controlCmd.AddCommand(s3control_getBucketVersioningCmd)
}
