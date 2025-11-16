package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_putBucketLifecycleConfigurationCmd = &cobra.Command{
	Use:   "put-bucket-lifecycle-configuration",
	Short: "This action puts a lifecycle configuration to an Amazon S3 on Outposts bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_putBucketLifecycleConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3control_putBucketLifecycleConfigurationCmd).Standalone()

		s3control_putBucketLifecycleConfigurationCmd.Flags().String("account-id", "", "The Amazon Web Services account ID of the Outposts bucket.")
		s3control_putBucketLifecycleConfigurationCmd.Flags().String("bucket", "", "The name of the bucket for which to set the configuration.")
		s3control_putBucketLifecycleConfigurationCmd.Flags().String("lifecycle-configuration", "", "Container for lifecycle rules.")
		s3control_putBucketLifecycleConfigurationCmd.MarkFlagRequired("account-id")
		s3control_putBucketLifecycleConfigurationCmd.MarkFlagRequired("bucket")
	})
	s3controlCmd.AddCommand(s3control_putBucketLifecycleConfigurationCmd)
}
