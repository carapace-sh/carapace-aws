package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ioteventsData_batchPutMessageCmd = &cobra.Command{
	Use:   "batch-put-message",
	Short: "Sends a set of messages to the IoT Events system.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ioteventsData_batchPutMessageCmd).Standalone()

	ioteventsData_batchPutMessageCmd.Flags().String("messages", "", "The list of messages to send.")
	ioteventsData_batchPutMessageCmd.MarkFlagRequired("messages")
	ioteventsDataCmd.AddCommand(ioteventsData_batchPutMessageCmd)
}
