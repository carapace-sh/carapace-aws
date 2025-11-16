package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmContacts_createContactChannelCmd = &cobra.Command{
	Use:   "create-contact-channel",
	Short: "A contact channel is the method that Incident Manager uses to engage your contact.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmContacts_createContactChannelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssmContacts_createContactChannelCmd).Standalone()

		ssmContacts_createContactChannelCmd.Flags().String("contact-id", "", "The Amazon Resource Name (ARN) of the contact you are adding the contact channel to.")
		ssmContacts_createContactChannelCmd.Flags().String("defer-activation", "", "If you want to activate the channel at a later time, you can choose to defer activation.")
		ssmContacts_createContactChannelCmd.Flags().String("delivery-address", "", "The details that Incident Manager uses when trying to engage the contact channel.")
		ssmContacts_createContactChannelCmd.Flags().String("idempotency-token", "", "A token ensuring that the operation is called only once with the specified details.")
		ssmContacts_createContactChannelCmd.Flags().String("name", "", "The name of the contact channel.")
		ssmContacts_createContactChannelCmd.Flags().String("type", "", "Incident Manager supports three types of contact channels:")
		ssmContacts_createContactChannelCmd.MarkFlagRequired("contact-id")
		ssmContacts_createContactChannelCmd.MarkFlagRequired("delivery-address")
		ssmContacts_createContactChannelCmd.MarkFlagRequired("name")
		ssmContacts_createContactChannelCmd.MarkFlagRequired("type")
	})
	ssmContactsCmd.AddCommand(ssmContacts_createContactChannelCmd)
}
