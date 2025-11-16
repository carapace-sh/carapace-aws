package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_putBucketAccelerateConfigurationCmd = &cobra.Command{
	Use:   "put-bucket-accelerate-configuration",
	Short: "This operation is not supported for directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_putBucketAccelerateConfigurationCmd).Standalone()

	s3_putBucketAccelerateConfigurationCmd.Flags().String("accelerate-configuration", "", "Container for setting the transfer acceleration state.")
	s3_putBucketAccelerateConfigurationCmd.Flags().String("bucket", "", "The name of the bucket for which the accelerate configuration is set.")
	s3_putBucketAccelerateConfigurationCmd.Flags().String("checksum-algorithm", "", "Indicates the algorithm used to create the checksum for the request when you use the SDK.")
	s3_putBucketAccelerateConfigurationCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
	s3_putBucketAccelerateConfigurationCmd.MarkFlagRequired("accelerate-configuration")
	s3_putBucketAccelerateConfigurationCmd.MarkFlagRequired("bucket")
	s3Cmd.AddCommand(s3_putBucketAccelerateConfigurationCmd)
}
