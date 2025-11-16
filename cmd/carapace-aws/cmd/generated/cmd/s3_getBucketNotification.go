package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_getBucketNotificationCmd = &cobra.Command{
	Use:   "get-bucket-notification",
	Short: "This operation is not supported for directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_getBucketNotificationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3_getBucketNotificationCmd).Standalone()

		s3_getBucketNotificationCmd.Flags().String("bucket", "", "The name of the bucket for which to get the notification configuration.")
		s3_getBucketNotificationCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
		s3_getBucketNotificationCmd.MarkFlagRequired("bucket")
	})
	s3Cmd.AddCommand(s3_getBucketNotificationCmd)
}
