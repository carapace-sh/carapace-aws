package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_createUploadCmd = &cobra.Command{
	Use:   "create-upload",
	Short: "Uploads an app or test scripts.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_createUploadCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(devicefarm_createUploadCmd).Standalone()

		devicefarm_createUploadCmd.Flags().String("content-type", "", "The upload's content type (for example, `application/octet-stream`).")
		devicefarm_createUploadCmd.Flags().String("name", "", "The upload's file name.")
		devicefarm_createUploadCmd.Flags().String("project-arn", "", "The ARN of the project for the upload.")
		devicefarm_createUploadCmd.Flags().String("type", "", "The upload's upload type.")
		devicefarm_createUploadCmd.MarkFlagRequired("name")
		devicefarm_createUploadCmd.MarkFlagRequired("project-arn")
		devicefarm_createUploadCmd.MarkFlagRequired("type")
	})
	devicefarmCmd.AddCommand(devicefarm_createUploadCmd)
}
