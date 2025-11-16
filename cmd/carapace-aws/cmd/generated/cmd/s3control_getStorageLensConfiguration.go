package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_getStorageLensConfigurationCmd = &cobra.Command{
	Use:   "get-storage-lens-configuration",
	Short: "This operation is not supported by directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_getStorageLensConfigurationCmd).Standalone()

	s3control_getStorageLensConfigurationCmd.Flags().String("account-id", "", "The account ID of the requester.")
	s3control_getStorageLensConfigurationCmd.Flags().String("config-id", "", "The ID of the Amazon S3 Storage Lens configuration.")
	s3control_getStorageLensConfigurationCmd.MarkFlagRequired("account-id")
	s3control_getStorageLensConfigurationCmd.MarkFlagRequired("config-id")
	s3controlCmd.AddCommand(s3control_getStorageLensConfigurationCmd)
}
