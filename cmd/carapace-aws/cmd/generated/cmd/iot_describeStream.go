package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_describeStreamCmd = &cobra.Command{
	Use:   "describe-stream",
	Short: "Gets information about a stream.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_describeStreamCmd).Standalone()

	iot_describeStreamCmd.Flags().String("stream-id", "", "The stream ID.")
	iot_describeStreamCmd.MarkFlagRequired("stream-id")
	iotCmd.AddCommand(iot_describeStreamCmd)
}
