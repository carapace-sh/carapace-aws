package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_putBucketNotificationConfigurationCmd = &cobra.Command{
	Use:   "put-bucket-notification-configuration",
	Short: "This operation is not supported for directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_putBucketNotificationConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3_putBucketNotificationConfigurationCmd).Standalone()

		s3_putBucketNotificationConfigurationCmd.Flags().String("bucket", "", "The name of the bucket.")
		s3_putBucketNotificationConfigurationCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
		s3_putBucketNotificationConfigurationCmd.Flags().String("notification-configuration", "", "")
		s3_putBucketNotificationConfigurationCmd.Flags().String("skip-destination-validation", "", "Skips validation of Amazon SQS, Amazon SNS, and Lambda destinations.")
		s3_putBucketNotificationConfigurationCmd.MarkFlagRequired("bucket")
		s3_putBucketNotificationConfigurationCmd.MarkFlagRequired("notification-configuration")
	})
	s3Cmd.AddCommand(s3_putBucketNotificationConfigurationCmd)
}
