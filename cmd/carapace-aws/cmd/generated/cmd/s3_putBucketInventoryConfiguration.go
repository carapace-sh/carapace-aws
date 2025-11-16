package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_putBucketInventoryConfigurationCmd = &cobra.Command{
	Use:   "put-bucket-inventory-configuration",
	Short: "This operation is not supported for directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_putBucketInventoryConfigurationCmd).Standalone()

	s3_putBucketInventoryConfigurationCmd.Flags().String("bucket", "", "The name of the bucket where the inventory configuration will be stored.")
	s3_putBucketInventoryConfigurationCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
	s3_putBucketInventoryConfigurationCmd.Flags().String("id", "", "The ID used to identify the inventory configuration.")
	s3_putBucketInventoryConfigurationCmd.Flags().String("inventory-configuration", "", "Specifies the inventory configuration.")
	s3_putBucketInventoryConfigurationCmd.MarkFlagRequired("bucket")
	s3_putBucketInventoryConfigurationCmd.MarkFlagRequired("id")
	s3_putBucketInventoryConfigurationCmd.MarkFlagRequired("inventory-configuration")
	s3Cmd.AddCommand(s3_putBucketInventoryConfigurationCmd)
}
