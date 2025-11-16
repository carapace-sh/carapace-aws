package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_notifyWhenUploadedCmd = &cobra.Command{
	Use:   "notify-when-uploaded",
	Short: "Sends you notification through Amazon EventBridge when all files written to your file share have been uploaded to Amazon S3.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_notifyWhenUploadedCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(storagegateway_notifyWhenUploadedCmd).Standalone()

		storagegateway_notifyWhenUploadedCmd.Flags().String("file-share-arn", "", "")
		storagegateway_notifyWhenUploadedCmd.MarkFlagRequired("file-share-arn")
	})
	storagegatewayCmd.AddCommand(storagegateway_notifyWhenUploadedCmd)
}
