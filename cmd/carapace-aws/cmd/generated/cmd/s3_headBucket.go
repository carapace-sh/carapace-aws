package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_headBucketCmd = &cobra.Command{
	Use:   "head-bucket",
	Short: "You can use this operation to determine if a bucket exists and if you have permission to access it.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_headBucketCmd).Standalone()

	s3_headBucketCmd.Flags().String("bucket", "", "The bucket name.")
	s3_headBucketCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
	s3_headBucketCmd.MarkFlagRequired("bucket")
	s3Cmd.AddCommand(s3_headBucketCmd)
}
