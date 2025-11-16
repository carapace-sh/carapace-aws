package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_getUploadCmd = &cobra.Command{
	Use:   "get-upload",
	Short: "Gets information about an upload.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_getUploadCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(devicefarm_getUploadCmd).Standalone()

		devicefarm_getUploadCmd.Flags().String("arn", "", "The upload's ARN.")
		devicefarm_getUploadCmd.MarkFlagRequired("arn")
	})
	devicefarmCmd.AddCommand(devicefarm_getUploadCmd)
}
