package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_completeAttachedFileUploadCmd = &cobra.Command{
	Use:   "complete-attached-file-upload",
	Short: "Allows you to confirm that the attached file has been uploaded using the pre-signed URL provided in the StartAttachedFileUpload API.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_completeAttachedFileUploadCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_completeAttachedFileUploadCmd).Standalone()

		connect_completeAttachedFileUploadCmd.Flags().String("associated-resource-arn", "", "The resource to which the attached file is (being) uploaded to.")
		connect_completeAttachedFileUploadCmd.Flags().String("file-id", "", "The unique identifier of the attached file resource.")
		connect_completeAttachedFileUploadCmd.Flags().String("instance-id", "", "The unique identifier of the Amazon Connect instance.")
		connect_completeAttachedFileUploadCmd.MarkFlagRequired("associated-resource-arn")
		connect_completeAttachedFileUploadCmd.MarkFlagRequired("file-id")
		connect_completeAttachedFileUploadCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_completeAttachedFileUploadCmd)
}
