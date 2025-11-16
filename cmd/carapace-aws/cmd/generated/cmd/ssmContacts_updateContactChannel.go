package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmContacts_updateContactChannelCmd = &cobra.Command{
	Use:   "update-contact-channel",
	Short: "Updates a contact's contact channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmContacts_updateContactChannelCmd).Standalone()

	ssmContacts_updateContactChannelCmd.Flags().String("contact-channel-id", "", "The Amazon Resource Name (ARN) of the contact channel you want to update.")
	ssmContacts_updateContactChannelCmd.Flags().String("delivery-address", "", "The details that Incident Manager uses when trying to engage the contact channel.")
	ssmContacts_updateContactChannelCmd.Flags().String("name", "", "The name of the contact channel.")
	ssmContacts_updateContactChannelCmd.MarkFlagRequired("contact-channel-id")
	ssmContactsCmd.AddCommand(ssmContacts_updateContactChannelCmd)
}
