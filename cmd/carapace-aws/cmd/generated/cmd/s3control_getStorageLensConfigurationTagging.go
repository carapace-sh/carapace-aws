package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_getStorageLensConfigurationTaggingCmd = &cobra.Command{
	Use:   "get-storage-lens-configuration-tagging",
	Short: "This operation is not supported by directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_getStorageLensConfigurationTaggingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3control_getStorageLensConfigurationTaggingCmd).Standalone()

		s3control_getStorageLensConfigurationTaggingCmd.Flags().String("account-id", "", "The account ID of the requester.")
		s3control_getStorageLensConfigurationTaggingCmd.Flags().String("config-id", "", "The ID of the Amazon S3 Storage Lens configuration.")
		s3control_getStorageLensConfigurationTaggingCmd.MarkFlagRequired("account-id")
		s3control_getStorageLensConfigurationTaggingCmd.MarkFlagRequired("config-id")
	})
	s3controlCmd.AddCommand(s3control_getStorageLensConfigurationTaggingCmd)
}
