package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_getBucketCorsCmd = &cobra.Command{
	Use:   "get-bucket-cors",
	Short: "This operation is not supported for directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_getBucketCorsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3_getBucketCorsCmd).Standalone()

		s3_getBucketCorsCmd.Flags().String("bucket", "", "The bucket name for which to get the cors configuration.")
		s3_getBucketCorsCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
		s3_getBucketCorsCmd.MarkFlagRequired("bucket")
	})
	s3Cmd.AddCommand(s3_getBucketCorsCmd)
}
