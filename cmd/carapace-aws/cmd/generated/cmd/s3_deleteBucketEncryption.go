package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_deleteBucketEncryptionCmd = &cobra.Command{
	Use:   "delete-bucket-encryption",
	Short: "This implementation of the DELETE action resets the default encryption for the bucket as server-side encryption with Amazon S3 managed keys (SSE-S3).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_deleteBucketEncryptionCmd).Standalone()

	s3_deleteBucketEncryptionCmd.Flags().String("bucket", "", "The name of the bucket containing the server-side encryption configuration to delete.")
	s3_deleteBucketEncryptionCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
	s3_deleteBucketEncryptionCmd.MarkFlagRequired("bucket")
	s3Cmd.AddCommand(s3_deleteBucketEncryptionCmd)
}
