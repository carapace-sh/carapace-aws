package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_updateBucketMetadataInventoryTableConfigurationCmd = &cobra.Command{
	Use:   "update-bucket-metadata-inventory-table-configuration",
	Short: "Enables or disables a live inventory table for an S3 Metadata configuration on a general purpose bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_updateBucketMetadataInventoryTableConfigurationCmd).Standalone()

	s3_updateBucketMetadataInventoryTableConfigurationCmd.Flags().String("bucket", "", "The general purpose bucket that corresponds to the metadata configuration that you want to enable or disable an inventory table for.")
	s3_updateBucketMetadataInventoryTableConfigurationCmd.Flags().String("checksum-algorithm", "", "The checksum algorithm to use with your inventory table configuration.")
	s3_updateBucketMetadataInventoryTableConfigurationCmd.Flags().String("content-md5", "", "The `Content-MD5` header for the inventory table configuration.")
	s3_updateBucketMetadataInventoryTableConfigurationCmd.Flags().String("expected-bucket-owner", "", "The expected owner of the general purpose bucket that corresponds to the metadata table configuration that you want to enable or disable an inventory table for.")
	s3_updateBucketMetadataInventoryTableConfigurationCmd.Flags().String("inventory-table-configuration", "", "The contents of your inventory table configuration.")
	s3_updateBucketMetadataInventoryTableConfigurationCmd.MarkFlagRequired("bucket")
	s3_updateBucketMetadataInventoryTableConfigurationCmd.MarkFlagRequired("inventory-table-configuration")
	s3Cmd.AddCommand(s3_updateBucketMetadataInventoryTableConfigurationCmd)
}
