package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_batchGetAttachedFileMetadataCmd = &cobra.Command{
	Use:   "batch-get-attached-file-metadata",
	Short: "Allows you to retrieve metadata about multiple attached files on an associated resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_batchGetAttachedFileMetadataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_batchGetAttachedFileMetadataCmd).Standalone()

		connect_batchGetAttachedFileMetadataCmd.Flags().String("associated-resource-arn", "", "The resource to which the attached file is (being) uploaded to.")
		connect_batchGetAttachedFileMetadataCmd.Flags().String("file-ids", "", "The unique identifiers of the attached file resource.")
		connect_batchGetAttachedFileMetadataCmd.Flags().String("instance-id", "", "The unique identifier of the Connect instance.")
		connect_batchGetAttachedFileMetadataCmd.MarkFlagRequired("associated-resource-arn")
		connect_batchGetAttachedFileMetadataCmd.MarkFlagRequired("file-ids")
		connect_batchGetAttachedFileMetadataCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_batchGetAttachedFileMetadataCmd)
}
