package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_deleteStorageLensConfigurationTaggingCmd = &cobra.Command{
	Use:   "delete-storage-lens-configuration-tagging",
	Short: "This operation is not supported by directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_deleteStorageLensConfigurationTaggingCmd).Standalone()

	s3control_deleteStorageLensConfigurationTaggingCmd.Flags().String("account-id", "", "The account ID of the requester.")
	s3control_deleteStorageLensConfigurationTaggingCmd.Flags().String("config-id", "", "The ID of the S3 Storage Lens configuration.")
	s3control_deleteStorageLensConfigurationTaggingCmd.MarkFlagRequired("account-id")
	s3control_deleteStorageLensConfigurationTaggingCmd.MarkFlagRequired("config-id")
	s3controlCmd.AddCommand(s3control_deleteStorageLensConfigurationTaggingCmd)
}
