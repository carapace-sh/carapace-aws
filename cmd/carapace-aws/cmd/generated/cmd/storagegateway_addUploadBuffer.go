package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_addUploadBufferCmd = &cobra.Command{
	Use:   "add-upload-buffer",
	Short: "Configures one or more gateway local disks as upload buffer for a specified gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_addUploadBufferCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(storagegateway_addUploadBufferCmd).Standalone()

		storagegateway_addUploadBufferCmd.Flags().String("disk-ids", "", "An array of strings that identify disks that are to be configured as working storage.")
		storagegateway_addUploadBufferCmd.Flags().String("gateway-arn", "", "")
		storagegateway_addUploadBufferCmd.MarkFlagRequired("disk-ids")
		storagegateway_addUploadBufferCmd.MarkFlagRequired("gateway-arn")
	})
	storagegatewayCmd.AddCommand(storagegateway_addUploadBufferCmd)
}
