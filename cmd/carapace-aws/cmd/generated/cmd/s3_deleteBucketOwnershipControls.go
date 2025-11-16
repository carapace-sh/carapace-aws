package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_deleteBucketOwnershipControlsCmd = &cobra.Command{
	Use:   "delete-bucket-ownership-controls",
	Short: "This operation is not supported for directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_deleteBucketOwnershipControlsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3_deleteBucketOwnershipControlsCmd).Standalone()

		s3_deleteBucketOwnershipControlsCmd.Flags().String("bucket", "", "The Amazon S3 bucket whose `OwnershipControls` you want to delete.")
		s3_deleteBucketOwnershipControlsCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
		s3_deleteBucketOwnershipControlsCmd.MarkFlagRequired("bucket")
	})
	s3Cmd.AddCommand(s3_deleteBucketOwnershipControlsCmd)
}
