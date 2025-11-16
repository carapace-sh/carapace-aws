package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_createStreamCmd = &cobra.Command{
	Use:   "create-stream",
	Short: "Creates a stream for delivering one or more large files in chunks over MQTT.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_createStreamCmd).Standalone()

	iot_createStreamCmd.Flags().String("description", "", "A description of the stream.")
	iot_createStreamCmd.Flags().String("files", "", "The files to stream.")
	iot_createStreamCmd.Flags().String("role-arn", "", "An IAM role that allows the IoT service principal to access your S3 files.")
	iot_createStreamCmd.Flags().String("stream-id", "", "The stream ID.")
	iot_createStreamCmd.Flags().String("tags", "", "Metadata which can be used to manage streams.")
	iot_createStreamCmd.MarkFlagRequired("files")
	iot_createStreamCmd.MarkFlagRequired("role-arn")
	iot_createStreamCmd.MarkFlagRequired("stream-id")
	iotCmd.AddCommand(iot_createStreamCmd)
}
