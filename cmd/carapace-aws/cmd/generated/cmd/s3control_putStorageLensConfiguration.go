package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_putStorageLensConfigurationCmd = &cobra.Command{
	Use:   "put-storage-lens-configuration",
	Short: "This operation is not supported by directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_putStorageLensConfigurationCmd).Standalone()

	s3control_putStorageLensConfigurationCmd.Flags().String("account-id", "", "The account ID of the requester.")
	s3control_putStorageLensConfigurationCmd.Flags().String("config-id", "", "The ID of the S3 Storage Lens configuration.")
	s3control_putStorageLensConfigurationCmd.Flags().String("storage-lens-configuration", "", "The S3 Storage Lens configuration.")
	s3control_putStorageLensConfigurationCmd.Flags().String("tags", "", "The tag set of the S3 Storage Lens configuration.")
	s3control_putStorageLensConfigurationCmd.MarkFlagRequired("account-id")
	s3control_putStorageLensConfigurationCmd.MarkFlagRequired("config-id")
	s3control_putStorageLensConfigurationCmd.MarkFlagRequired("storage-lens-configuration")
	s3controlCmd.AddCommand(s3control_putStorageLensConfigurationCmd)
}
