package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_deleteStreamCmd = &cobra.Command{
	Use:   "delete-stream",
	Short: "Deletes a stream.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_deleteStreamCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_deleteStreamCmd).Standalone()

		iot_deleteStreamCmd.Flags().String("stream-id", "", "The stream ID.")
		iot_deleteStreamCmd.MarkFlagRequired("stream-id")
	})
	iotCmd.AddCommand(iot_deleteStreamCmd)
}
