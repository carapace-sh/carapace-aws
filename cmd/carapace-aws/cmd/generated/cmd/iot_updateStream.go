package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_updateStreamCmd = &cobra.Command{
	Use:   "update-stream",
	Short: "Updates an existing stream.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_updateStreamCmd).Standalone()

	iot_updateStreamCmd.Flags().String("description", "", "The description of the stream.")
	iot_updateStreamCmd.Flags().String("files", "", "The files associated with the stream.")
	iot_updateStreamCmd.Flags().String("role-arn", "", "An IAM role that allows the IoT service principal assumes to access your S3 files.")
	iot_updateStreamCmd.Flags().String("stream-id", "", "The stream ID.")
	iot_updateStreamCmd.MarkFlagRequired("stream-id")
	iotCmd.AddCommand(iot_updateStreamCmd)
}
