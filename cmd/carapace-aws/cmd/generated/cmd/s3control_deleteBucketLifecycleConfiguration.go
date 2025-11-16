package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_deleteBucketLifecycleConfigurationCmd = &cobra.Command{
	Use:   "delete-bucket-lifecycle-configuration",
	Short: "This action deletes an Amazon S3 on Outposts bucket's lifecycle configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_deleteBucketLifecycleConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3control_deleteBucketLifecycleConfigurationCmd).Standalone()

		s3control_deleteBucketLifecycleConfigurationCmd.Flags().String("account-id", "", "The account ID of the lifecycle configuration to delete.")
		s3control_deleteBucketLifecycleConfigurationCmd.Flags().String("bucket", "", "Specifies the bucket.")
		s3control_deleteBucketLifecycleConfigurationCmd.MarkFlagRequired("account-id")
		s3control_deleteBucketLifecycleConfigurationCmd.MarkFlagRequired("bucket")
	})
	s3controlCmd.AddCommand(s3control_deleteBucketLifecycleConfigurationCmd)
}
