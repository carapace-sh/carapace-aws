package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_deleteUploadCmd = &cobra.Command{
	Use:   "delete-upload",
	Short: "Deletes an upload given the upload ARN.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_deleteUploadCmd).Standalone()

	devicefarm_deleteUploadCmd.Flags().String("arn", "", "Represents the Amazon Resource Name (ARN) of the Device Farm upload to delete.")
	devicefarm_deleteUploadCmd.MarkFlagRequired("arn")
	devicefarmCmd.AddCommand(devicefarm_deleteUploadCmd)
}
