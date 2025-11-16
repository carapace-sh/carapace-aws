package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_getBucketEncryptionCmd = &cobra.Command{
	Use:   "get-bucket-encryption",
	Short: "Returns the default encryption configuration for an Amazon S3 bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_getBucketEncryptionCmd).Standalone()

	s3_getBucketEncryptionCmd.Flags().String("bucket", "", "The name of the bucket from which the server-side encryption configuration is retrieved.")
	s3_getBucketEncryptionCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
	s3_getBucketEncryptionCmd.MarkFlagRequired("bucket")
	s3Cmd.AddCommand(s3_getBucketEncryptionCmd)
}
