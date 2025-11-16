package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_deleteAttachedFileCmd = &cobra.Command{
	Use:   "delete-attached-file",
	Short: "Deletes an attached file along with the underlying S3 Object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_deleteAttachedFileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_deleteAttachedFileCmd).Standalone()

		connect_deleteAttachedFileCmd.Flags().String("associated-resource-arn", "", "The resource to which the attached file is (being) uploaded to.")
		connect_deleteAttachedFileCmd.Flags().String("file-id", "", "The unique identifier of the attached file resource.")
		connect_deleteAttachedFileCmd.Flags().String("instance-id", "", "The unique identifier of the Connect instance.")
		connect_deleteAttachedFileCmd.MarkFlagRequired("associated-resource-arn")
		connect_deleteAttachedFileCmd.MarkFlagRequired("file-id")
		connect_deleteAttachedFileCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_deleteAttachedFileCmd)
}
