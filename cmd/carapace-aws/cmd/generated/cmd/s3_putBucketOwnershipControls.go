package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_putBucketOwnershipControlsCmd = &cobra.Command{
	Use:   "put-bucket-ownership-controls",
	Short: "This operation is not supported for directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_putBucketOwnershipControlsCmd).Standalone()

	s3_putBucketOwnershipControlsCmd.Flags().String("bucket", "", "The name of the Amazon S3 bucket whose `OwnershipControls` you want to set.")
	s3_putBucketOwnershipControlsCmd.Flags().String("checksum-algorithm", "", "Indicates the algorithm used to create the checksum for the object when you use the SDK.")
	s3_putBucketOwnershipControlsCmd.Flags().String("content-md5", "", "The MD5 hash of the `OwnershipControls` request body.")
	s3_putBucketOwnershipControlsCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
	s3_putBucketOwnershipControlsCmd.Flags().String("ownership-controls", "", "The `OwnershipControls` (BucketOwnerEnforced, BucketOwnerPreferred, or ObjectWriter) that you want to apply to this Amazon S3 bucket.")
	s3_putBucketOwnershipControlsCmd.MarkFlagRequired("bucket")
	s3_putBucketOwnershipControlsCmd.MarkFlagRequired("ownership-controls")
	s3Cmd.AddCommand(s3_putBucketOwnershipControlsCmd)
}
