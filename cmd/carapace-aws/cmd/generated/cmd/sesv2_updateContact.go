package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_updateContactCmd = &cobra.Command{
	Use:   "update-contact",
	Short: "Updates a contact's preferences for a list.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_updateContactCmd).Standalone()

	sesv2_updateContactCmd.Flags().String("attributes-data", "", "The attribute data attached to a contact.")
	sesv2_updateContactCmd.Flags().String("contact-list-name", "", "The name of the contact list.")
	sesv2_updateContactCmd.Flags().String("email-address", "", "The contact's email address.")
	sesv2_updateContactCmd.Flags().String("topic-preferences", "", "The contact's preference for being opted-in to or opted-out of a topic.")
	sesv2_updateContactCmd.Flags().String("unsubscribe-all", "", "A boolean value status noting if the contact is unsubscribed from all contact list topics.")
	sesv2_updateContactCmd.MarkFlagRequired("contact-list-name")
	sesv2_updateContactCmd.MarkFlagRequired("email-address")
	sesv2Cmd.AddCommand(sesv2_updateContactCmd)
}
