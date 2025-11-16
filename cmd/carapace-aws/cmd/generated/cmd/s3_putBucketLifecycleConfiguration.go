package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_putBucketLifecycleConfigurationCmd = &cobra.Command{
	Use:   "put-bucket-lifecycle-configuration",
	Short: "Creates a new lifecycle configuration for the bucket or replaces an existing lifecycle configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_putBucketLifecycleConfigurationCmd).Standalone()

	s3_putBucketLifecycleConfigurationCmd.Flags().String("bucket", "", "The name of the bucket for which to set the configuration.")
	s3_putBucketLifecycleConfigurationCmd.Flags().String("checksum-algorithm", "", "Indicates the algorithm used to create the checksum for the request when you use the SDK.")
	s3_putBucketLifecycleConfigurationCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
	s3_putBucketLifecycleConfigurationCmd.Flags().String("lifecycle-configuration", "", "Container for lifecycle rules.")
	s3_putBucketLifecycleConfigurationCmd.Flags().String("transition-default-minimum-object-size", "", "Indicates which default minimum object size behavior is applied to the lifecycle configuration.")
	s3_putBucketLifecycleConfigurationCmd.MarkFlagRequired("bucket")
	s3Cmd.AddCommand(s3_putBucketLifecycleConfigurationCmd)
}
