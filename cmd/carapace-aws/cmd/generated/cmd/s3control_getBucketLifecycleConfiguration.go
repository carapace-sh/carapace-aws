package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_getBucketLifecycleConfigurationCmd = &cobra.Command{
	Use:   "get-bucket-lifecycle-configuration",
	Short: "This action gets an Amazon S3 on Outposts bucket's lifecycle configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_getBucketLifecycleConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3control_getBucketLifecycleConfigurationCmd).Standalone()

		s3control_getBucketLifecycleConfigurationCmd.Flags().String("account-id", "", "The Amazon Web Services account ID of the Outposts bucket.")
		s3control_getBucketLifecycleConfigurationCmd.Flags().String("bucket", "", "The Amazon Resource Name (ARN) of the bucket.")
		s3control_getBucketLifecycleConfigurationCmd.MarkFlagRequired("account-id")
		s3control_getBucketLifecycleConfigurationCmd.MarkFlagRequired("bucket")
	})
	s3controlCmd.AddCommand(s3control_getBucketLifecycleConfigurationCmd)
}
