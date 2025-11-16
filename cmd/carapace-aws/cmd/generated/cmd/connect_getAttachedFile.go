package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_getAttachedFileCmd = &cobra.Command{
	Use:   "get-attached-file",
	Short: "Provides a pre-signed URL for download of an approved attached file.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_getAttachedFileCmd).Standalone()

	connect_getAttachedFileCmd.Flags().String("associated-resource-arn", "", "The resource to which the attached file is (being) uploaded to.")
	connect_getAttachedFileCmd.Flags().String("file-id", "", "The unique identifier of the attached file resource.")
	connect_getAttachedFileCmd.Flags().String("instance-id", "", "The unique identifier of the Amazon Connect instance.")
	connect_getAttachedFileCmd.Flags().String("url-expiry-in-seconds", "", "Optional override for the expiry of the pre-signed S3 URL in seconds.")
	connect_getAttachedFileCmd.MarkFlagRequired("associated-resource-arn")
	connect_getAttachedFileCmd.MarkFlagRequired("file-id")
	connect_getAttachedFileCmd.MarkFlagRequired("instance-id")
	connectCmd.AddCommand(connect_getAttachedFileCmd)
}
