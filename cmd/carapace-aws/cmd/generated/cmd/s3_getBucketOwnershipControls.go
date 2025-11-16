package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_getBucketOwnershipControlsCmd = &cobra.Command{
	Use:   "get-bucket-ownership-controls",
	Short: "This operation is not supported for directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_getBucketOwnershipControlsCmd).Standalone()

	s3_getBucketOwnershipControlsCmd.Flags().String("bucket", "", "The name of the Amazon S3 bucket whose `OwnershipControls` you want to retrieve.")
	s3_getBucketOwnershipControlsCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
	s3_getBucketOwnershipControlsCmd.MarkFlagRequired("bucket")
	s3Cmd.AddCommand(s3_getBucketOwnershipControlsCmd)
}
