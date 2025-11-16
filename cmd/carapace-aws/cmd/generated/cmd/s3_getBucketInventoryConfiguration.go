package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_getBucketInventoryConfigurationCmd = &cobra.Command{
	Use:   "get-bucket-inventory-configuration",
	Short: "This operation is not supported for directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_getBucketInventoryConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3_getBucketInventoryConfigurationCmd).Standalone()

		s3_getBucketInventoryConfigurationCmd.Flags().String("bucket", "", "The name of the bucket containing the inventory configuration to retrieve.")
		s3_getBucketInventoryConfigurationCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
		s3_getBucketInventoryConfigurationCmd.Flags().String("id", "", "The ID used to identify the inventory configuration.")
		s3_getBucketInventoryConfigurationCmd.MarkFlagRequired("bucket")
		s3_getBucketInventoryConfigurationCmd.MarkFlagRequired("id")
	})
	s3Cmd.AddCommand(s3_getBucketInventoryConfigurationCmd)
}
