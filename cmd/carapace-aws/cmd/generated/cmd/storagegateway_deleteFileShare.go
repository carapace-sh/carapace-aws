package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_deleteFileShareCmd = &cobra.Command{
	Use:   "delete-file-share",
	Short: "Deletes a file share from an S3 File Gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_deleteFileShareCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(storagegateway_deleteFileShareCmd).Standalone()

		storagegateway_deleteFileShareCmd.Flags().String("file-share-arn", "", "The Amazon Resource Name (ARN) of the file share to be deleted.")
		storagegateway_deleteFileShareCmd.Flags().String("force-delete", "", "If this value is set to `true`, the operation deletes a file share immediately and aborts all data uploads to Amazon Web Services.")
		storagegateway_deleteFileShareCmd.MarkFlagRequired("file-share-arn")
	})
	storagegatewayCmd.AddCommand(storagegateway_deleteFileShareCmd)
}
