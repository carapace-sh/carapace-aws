package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_getBucketLifecycleConfigurationCmd = &cobra.Command{
	Use:   "get-bucket-lifecycle-configuration",
	Short: "Returns the lifecycle configuration information set on the bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_getBucketLifecycleConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3_getBucketLifecycleConfigurationCmd).Standalone()

		s3_getBucketLifecycleConfigurationCmd.Flags().String("bucket", "", "The name of the bucket for which to get the lifecycle information.")
		s3_getBucketLifecycleConfigurationCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
		s3_getBucketLifecycleConfigurationCmd.MarkFlagRequired("bucket")
	})
	s3Cmd.AddCommand(s3_getBucketLifecycleConfigurationCmd)
}
