package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_createContactCmd = &cobra.Command{
	Use:   "create-contact",
	Short: "Only the VOICE, EMAIL, and TASK channels are supported.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_createContactCmd).Standalone()

	connect_createContactCmd.Flags().String("attributes", "", "A custom key-value pair using an attribute map.")
	connect_createContactCmd.Flags().String("channel", "", "The channel for the contact.")
	connect_createContactCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	connect_createContactCmd.Flags().String("description", "", "A description of the contact.")
	connect_createContactCmd.Flags().String("expiry-duration-in-minutes", "", "Number of minutes the contact will be active for before expiring")
	connect_createContactCmd.Flags().String("initiate-as", "", "Initial state of the contact when it's created.")
	connect_createContactCmd.Flags().String("initiation-method", "", "Indicates how the contact was initiated.")
	connect_createContactCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_createContactCmd.Flags().String("name", "", "The name of a the contact.")
	connect_createContactCmd.Flags().String("previous-contact-id", "", "The ID of the previous contact when creating a transfer contact.")
	connect_createContactCmd.Flags().String("references", "", "A formatted URL that is shown to an agent in the Contact Control Panel (CCP).")
	connect_createContactCmd.Flags().String("related-contact-id", "", "The identifier of the contact in this instance of Amazon Connect.")
	connect_createContactCmd.Flags().String("segment-attributes", "", "A set of system defined key-value pairs stored on individual contact segments (unique contact ID) using an attribute map.")
	connect_createContactCmd.Flags().String("user-info", "", "User details for the contact")
	connect_createContactCmd.MarkFlagRequired("channel")
	connect_createContactCmd.MarkFlagRequired("initiation-method")
	connect_createContactCmd.MarkFlagRequired("instance-id")
	connectCmd.AddCommand(connect_createContactCmd)
}
