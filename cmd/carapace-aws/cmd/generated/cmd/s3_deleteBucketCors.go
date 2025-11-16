package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_deleteBucketCorsCmd = &cobra.Command{
	Use:   "delete-bucket-cors",
	Short: "This operation is not supported for directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_deleteBucketCorsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3_deleteBucketCorsCmd).Standalone()

		s3_deleteBucketCorsCmd.Flags().String("bucket", "", "Specifies the bucket whose `cors` configuration is being deleted.")
		s3_deleteBucketCorsCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
		s3_deleteBucketCorsCmd.MarkFlagRequired("bucket")
	})
	s3Cmd.AddCommand(s3_deleteBucketCorsCmd)
}
