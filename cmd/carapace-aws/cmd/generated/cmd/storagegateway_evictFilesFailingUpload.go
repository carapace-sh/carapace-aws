package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_evictFilesFailingUploadCmd = &cobra.Command{
	Use:   "evict-files-failing-upload",
	Short: "Starts a process that cleans the specified file share's cache of file entries that are failing upload to Amazon S3.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_evictFilesFailingUploadCmd).Standalone()

	storagegateway_evictFilesFailingUploadCmd.Flags().String("file-share-arn", "", "The Amazon Resource Name (ARN) of the file share for which you want to start the cache clean operation.")
	storagegateway_evictFilesFailingUploadCmd.Flags().String("force-remove", "", "Specifies whether cache entries with full or partial file data currently stored on the gateway will be forcibly removed by the cache clean operation.")
	storagegateway_evictFilesFailingUploadCmd.MarkFlagRequired("file-share-arn")
	storagegatewayCmd.AddCommand(storagegateway_evictFilesFailingUploadCmd)
}
