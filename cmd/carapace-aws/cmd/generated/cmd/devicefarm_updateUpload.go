package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_updateUploadCmd = &cobra.Command{
	Use:   "update-upload",
	Short: "Updates an uploaded test spec.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_updateUploadCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(devicefarm_updateUploadCmd).Standalone()

		devicefarm_updateUploadCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the uploaded test spec.")
		devicefarm_updateUploadCmd.Flags().String("content-type", "", "The upload's content type (for example, `application/x-yaml`).")
		devicefarm_updateUploadCmd.Flags().Bool("edit-content", false, "Set to true if the YAML file has changed and must be updated.")
		devicefarm_updateUploadCmd.Flags().String("name", "", "The upload's test spec file name.")
		devicefarm_updateUploadCmd.Flags().Bool("no-edit-content", false, "Set to true if the YAML file has changed and must be updated.")
		devicefarm_updateUploadCmd.MarkFlagRequired("arn")
		devicefarm_updateUploadCmd.Flag("no-edit-content").Hidden = true
	})
	devicefarmCmd.AddCommand(devicefarm_updateUploadCmd)
}
