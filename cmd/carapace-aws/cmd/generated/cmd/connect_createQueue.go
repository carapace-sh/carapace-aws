package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_createQueueCmd = &cobra.Command{
	Use:   "create-queue",
	Short: "Creates a new queue for the specified Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_createQueueCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_createQueueCmd).Standalone()

		connect_createQueueCmd.Flags().String("description", "", "The description of the queue.")
		connect_createQueueCmd.Flags().String("hours-of-operation-id", "", "The identifier for the hours of operation.")
		connect_createQueueCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_createQueueCmd.Flags().String("max-contacts", "", "The maximum number of contacts that can be in the queue before it is considered full.")
		connect_createQueueCmd.Flags().String("name", "", "The name of the queue.")
		connect_createQueueCmd.Flags().String("outbound-caller-config", "", "The outbound caller ID name, number, and outbound whisper flow.")
		connect_createQueueCmd.Flags().String("outbound-email-config", "", "The outbound email address ID for a specified queue.")
		connect_createQueueCmd.Flags().String("quick-connect-ids", "", "The quick connects available to agents who are working the queue.")
		connect_createQueueCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for this resource.")
		connect_createQueueCmd.MarkFlagRequired("hours-of-operation-id")
		connect_createQueueCmd.MarkFlagRequired("instance-id")
		connect_createQueueCmd.MarkFlagRequired("name")
	})
	connectCmd.AddCommand(connect_createQueueCmd)
}
