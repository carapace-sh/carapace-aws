package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_startAttachedFileUploadCmd = &cobra.Command{
	Use:   "start-attached-file-upload",
	Short: "Provides a pre-signed Amazon S3 URL in response for uploading your content.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_startAttachedFileUploadCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_startAttachedFileUploadCmd).Standalone()

		connect_startAttachedFileUploadCmd.Flags().String("associated-resource-arn", "", "The resource to which the attached file is (being) uploaded to.")
		connect_startAttachedFileUploadCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		connect_startAttachedFileUploadCmd.Flags().String("created-by", "", "Represents the identity that created the file.")
		connect_startAttachedFileUploadCmd.Flags().String("file-name", "", "A case-sensitive name of the attached file being uploaded.")
		connect_startAttachedFileUploadCmd.Flags().String("file-size-in-bytes", "", "The size of the attached file in bytes.")
		connect_startAttachedFileUploadCmd.Flags().String("file-use-case-type", "", "The use case for the file.")
		connect_startAttachedFileUploadCmd.Flags().String("instance-id", "", "The unique identifier of the Amazon Connect instance.")
		connect_startAttachedFileUploadCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for this resource.")
		connect_startAttachedFileUploadCmd.Flags().String("url-expiry-in-seconds", "", "Optional override for the expiry of the pre-signed S3 URL in seconds.")
		connect_startAttachedFileUploadCmd.MarkFlagRequired("associated-resource-arn")
		connect_startAttachedFileUploadCmd.MarkFlagRequired("file-name")
		connect_startAttachedFileUploadCmd.MarkFlagRequired("file-size-in-bytes")
		connect_startAttachedFileUploadCmd.MarkFlagRequired("file-use-case-type")
		connect_startAttachedFileUploadCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_startAttachedFileUploadCmd)
}
