package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_getBucketAccelerateConfigurationCmd = &cobra.Command{
	Use:   "get-bucket-accelerate-configuration",
	Short: "This operation is not supported for directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_getBucketAccelerateConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3_getBucketAccelerateConfigurationCmd).Standalone()

		s3_getBucketAccelerateConfigurationCmd.Flags().String("bucket", "", "The name of the bucket for which the accelerate configuration is retrieved.")
		s3_getBucketAccelerateConfigurationCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
		s3_getBucketAccelerateConfigurationCmd.Flags().String("request-payer", "", "")
		s3_getBucketAccelerateConfigurationCmd.MarkFlagRequired("bucket")
	})
	s3Cmd.AddCommand(s3_getBucketAccelerateConfigurationCmd)
}
