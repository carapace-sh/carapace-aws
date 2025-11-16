package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_deleteBucketInventoryConfigurationCmd = &cobra.Command{
	Use:   "delete-bucket-inventory-configuration",
	Short: "This operation is not supported for directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_deleteBucketInventoryConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3_deleteBucketInventoryConfigurationCmd).Standalone()

		s3_deleteBucketInventoryConfigurationCmd.Flags().String("bucket", "", "The name of the bucket containing the inventory configuration to delete.")
		s3_deleteBucketInventoryConfigurationCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
		s3_deleteBucketInventoryConfigurationCmd.Flags().String("id", "", "The ID used to identify the inventory configuration.")
		s3_deleteBucketInventoryConfigurationCmd.MarkFlagRequired("bucket")
		s3_deleteBucketInventoryConfigurationCmd.MarkFlagRequired("id")
	})
	s3Cmd.AddCommand(s3_deleteBucketInventoryConfigurationCmd)
}
