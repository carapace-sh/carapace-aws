package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotthingsgraph_getUploadStatusCmd = &cobra.Command{
	Use:   "get-upload-status",
	Short: "Gets the status of the specified upload.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotthingsgraph_getUploadStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotthingsgraph_getUploadStatusCmd).Standalone()

		iotthingsgraph_getUploadStatusCmd.Flags().String("upload-id", "", "The ID of the upload.")
		iotthingsgraph_getUploadStatusCmd.MarkFlagRequired("upload-id")
	})
	iotthingsgraphCmd.AddCommand(iotthingsgraph_getUploadStatusCmd)
}
