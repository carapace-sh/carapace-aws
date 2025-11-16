package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_deleteBucketCmd = &cobra.Command{
	Use:   "delete-bucket",
	Short: "Deletes the S3 bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_deleteBucketCmd).Standalone()

	s3_deleteBucketCmd.Flags().String("bucket", "", "Specifies the bucket being deleted.")
	s3_deleteBucketCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
	s3_deleteBucketCmd.MarkFlagRequired("bucket")
	s3Cmd.AddCommand(s3_deleteBucketCmd)
}
