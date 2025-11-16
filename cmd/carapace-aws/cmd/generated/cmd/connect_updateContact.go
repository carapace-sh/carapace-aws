package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_updateContactCmd = &cobra.Command{
	Use:   "update-contact",
	Short: "This API is in preview release for Amazon Connect and is subject to change.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_updateContactCmd).Standalone()

	connect_updateContactCmd.Flags().String("contact-id", "", "The identifier of the contact.")
	connect_updateContactCmd.Flags().String("customer-endpoint", "", "The endpoint of the customer for which the contact was initiated.")
	connect_updateContactCmd.Flags().String("description", "", "The description of the contact.")
	connect_updateContactCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_updateContactCmd.Flags().String("name", "", "The name of the contact.")
	connect_updateContactCmd.Flags().String("queue-info", "", "Information about the queue associated with a contact.")
	connect_updateContactCmd.Flags().String("references", "", "Well-formed data on contact, shown to agents on Contact Control Panel (CCP).")
	connect_updateContactCmd.Flags().String("segment-attributes", "", "A set of system defined key-value pairs stored on individual contact segments (unique contact ID) using an attribute map.")
	connect_updateContactCmd.Flags().String("system-endpoint", "", "External system endpoint for the contact was initiated.")
	connect_updateContactCmd.Flags().String("user-info", "", "Information about the agent associated with a contact.")
	connect_updateContactCmd.MarkFlagRequired("contact-id")
	connect_updateContactCmd.MarkFlagRequired("instance-id")
	connectCmd.AddCommand(connect_updateContactCmd)
}
