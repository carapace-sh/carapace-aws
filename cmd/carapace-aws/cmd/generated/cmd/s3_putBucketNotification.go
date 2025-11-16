package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_putBucketNotificationCmd = &cobra.Command{
	Use:   "put-bucket-notification",
	Short: "This operation is not supported for directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_putBucketNotificationCmd).Standalone()

	s3_putBucketNotificationCmd.Flags().String("bucket", "", "The name of the bucket.")
	s3_putBucketNotificationCmd.Flags().String("checksum-algorithm", "", "Indicates the algorithm used to create the checksum for the request when you use the SDK.")
	s3_putBucketNotificationCmd.Flags().String("content-md5", "", "The MD5 hash of the `PutPublicAccessBlock` request body.")
	s3_putBucketNotificationCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
	s3_putBucketNotificationCmd.Flags().String("notification-configuration", "", "The container for the configuration.")
	s3_putBucketNotificationCmd.MarkFlagRequired("bucket")
	s3_putBucketNotificationCmd.MarkFlagRequired("notification-configuration")
	s3Cmd.AddCommand(s3_putBucketNotificationCmd)
}
