package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_putStorageLensConfigurationTaggingCmd = &cobra.Command{
	Use:   "put-storage-lens-configuration-tagging",
	Short: "This operation is not supported by directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_putStorageLensConfigurationTaggingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3control_putStorageLensConfigurationTaggingCmd).Standalone()

		s3control_putStorageLensConfigurationTaggingCmd.Flags().String("account-id", "", "The account ID of the requester.")
		s3control_putStorageLensConfigurationTaggingCmd.Flags().String("config-id", "", "The ID of the S3 Storage Lens configuration.")
		s3control_putStorageLensConfigurationTaggingCmd.Flags().String("tags", "", "The tag set of the S3 Storage Lens configuration.")
		s3control_putStorageLensConfigurationTaggingCmd.MarkFlagRequired("account-id")
		s3control_putStorageLensConfigurationTaggingCmd.MarkFlagRequired("config-id")
		s3control_putStorageLensConfigurationTaggingCmd.MarkFlagRequired("tags")
	})
	s3controlCmd.AddCommand(s3control_putStorageLensConfigurationTaggingCmd)
}
