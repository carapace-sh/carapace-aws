package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_deleteBucketLifecycleCmd = &cobra.Command{
	Use:   "delete-bucket-lifecycle",
	Short: "Deletes the lifecycle configuration from the specified bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_deleteBucketLifecycleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3_deleteBucketLifecycleCmd).Standalone()

		s3_deleteBucketLifecycleCmd.Flags().String("bucket", "", "The bucket name of the lifecycle to delete.")
		s3_deleteBucketLifecycleCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
		s3_deleteBucketLifecycleCmd.MarkFlagRequired("bucket")
	})
	s3Cmd.AddCommand(s3_deleteBucketLifecycleCmd)
}
