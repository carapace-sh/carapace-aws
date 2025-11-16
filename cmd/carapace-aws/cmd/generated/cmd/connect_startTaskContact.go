package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_startTaskContactCmd = &cobra.Command{
	Use:   "start-task-contact",
	Short: "Initiates a flow to start a new task contact.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_startTaskContactCmd).Standalone()

	connect_startTaskContactCmd.Flags().String("attributes", "", "A custom key-value pair using an attribute map.")
	connect_startTaskContactCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	connect_startTaskContactCmd.Flags().String("contact-flow-id", "", "The identifier of the flow for initiating the tasks.")
	connect_startTaskContactCmd.Flags().String("description", "", "A description of the task that is shown to an agent in the Contact Control Panel (CCP).")
	connect_startTaskContactCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_startTaskContactCmd.Flags().String("name", "", "The name of a task that is shown to an agent in the Contact Control Panel (CCP).")
	connect_startTaskContactCmd.Flags().String("previous-contact-id", "", "The identifier of the previous chat, voice, or task contact.")
	connect_startTaskContactCmd.Flags().String("quick-connect-id", "", "The identifier for the quick connect.")
	connect_startTaskContactCmd.Flags().String("references", "", "A formatted URL that is shown to an agent in the Contact Control Panel (CCP).")
	connect_startTaskContactCmd.Flags().String("related-contact-id", "", "The contactId that is [related](https://docs.aws.amazon.com/connect/latest/adminguide/tasks.html#linked-tasks) to this contact.")
	connect_startTaskContactCmd.Flags().String("scheduled-time", "", "The timestamp, in Unix Epoch seconds format, at which to start running the inbound flow.")
	connect_startTaskContactCmd.Flags().String("segment-attributes", "", "A set of system defined key-value pairs stored on individual contact segments (unique contact ID) using an attribute map.")
	connect_startTaskContactCmd.Flags().String("task-template-id", "", "A unique identifier for the task template.")
	connect_startTaskContactCmd.MarkFlagRequired("instance-id")
	connect_startTaskContactCmd.MarkFlagRequired("name")
	connectCmd.AddCommand(connect_startTaskContactCmd)
}
