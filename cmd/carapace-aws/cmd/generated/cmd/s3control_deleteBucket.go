package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_deleteBucketCmd = &cobra.Command{
	Use:   "delete-bucket",
	Short: "This action deletes an Amazon S3 on Outposts bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_deleteBucketCmd).Standalone()

	s3control_deleteBucketCmd.Flags().String("account-id", "", "The account ID that owns the Outposts bucket.")
	s3control_deleteBucketCmd.Flags().String("bucket", "", "Specifies the bucket being deleted.")
	s3control_deleteBucketCmd.MarkFlagRequired("account-id")
	s3control_deleteBucketCmd.MarkFlagRequired("bucket")
	s3controlCmd.AddCommand(s3control_deleteBucketCmd)
}
