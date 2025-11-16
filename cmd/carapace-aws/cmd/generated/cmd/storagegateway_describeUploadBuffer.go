package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_describeUploadBufferCmd = &cobra.Command{
	Use:   "describe-upload-buffer",
	Short: "Returns information about the upload buffer of a gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_describeUploadBufferCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(storagegateway_describeUploadBufferCmd).Standalone()

		storagegateway_describeUploadBufferCmd.Flags().String("gateway-arn", "", "")
		storagegateway_describeUploadBufferCmd.MarkFlagRequired("gateway-arn")
	})
	storagegatewayCmd.AddCommand(storagegateway_describeUploadBufferCmd)
}
