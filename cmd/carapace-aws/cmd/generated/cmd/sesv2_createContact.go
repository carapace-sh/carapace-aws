package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_createContactCmd = &cobra.Command{
	Use:   "create-contact",
	Short: "Creates a contact, which is an end-user who is receiving the email, and adds them to a contact list.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_createContactCmd).Standalone()

	sesv2_createContactCmd.Flags().String("attributes-data", "", "The attribute data attached to a contact.")
	sesv2_createContactCmd.Flags().String("contact-list-name", "", "The name of the contact list to which the contact should be added.")
	sesv2_createContactCmd.Flags().String("email-address", "", "The contact's email address.")
	sesv2_createContactCmd.Flags().String("topic-preferences", "", "The contact's preferences for being opted-in to or opted-out of topics.")
	sesv2_createContactCmd.Flags().String("unsubscribe-all", "", "A boolean value status noting if the contact is unsubscribed from all contact list topics.")
	sesv2_createContactCmd.MarkFlagRequired("contact-list-name")
	sesv2_createContactCmd.MarkFlagRequired("email-address")
	sesv2Cmd.AddCommand(sesv2_createContactCmd)
}
