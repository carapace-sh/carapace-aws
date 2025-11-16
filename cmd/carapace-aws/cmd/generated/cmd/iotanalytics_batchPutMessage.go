package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotanalytics_batchPutMessageCmd = &cobra.Command{
	Use:   "batch-put-message",
	Short: "Sends messages to a channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotanalytics_batchPutMessageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotanalytics_batchPutMessageCmd).Standalone()

		iotanalytics_batchPutMessageCmd.Flags().String("channel-name", "", "The name of the channel where the messages are sent.")
		iotanalytics_batchPutMessageCmd.Flags().String("messages", "", "The list of messages to be sent.")
		iotanalytics_batchPutMessageCmd.MarkFlagRequired("channel-name")
		iotanalytics_batchPutMessageCmd.MarkFlagRequired("messages")
	})
	iotanalyticsCmd.AddCommand(iotanalytics_batchPutMessageCmd)
}
