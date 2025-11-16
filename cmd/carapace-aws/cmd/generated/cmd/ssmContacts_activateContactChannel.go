package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmContacts_activateContactChannelCmd = &cobra.Command{
	Use:   "activate-contact-channel",
	Short: "Activates a contact's contact channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmContacts_activateContactChannelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssmContacts_activateContactChannelCmd).Standalone()

		ssmContacts_activateContactChannelCmd.Flags().String("activation-code", "", "The code sent to the contact channel when it was created in the contact.")
		ssmContacts_activateContactChannelCmd.Flags().String("contact-channel-id", "", "The Amazon Resource Name (ARN) of the contact channel.")
		ssmContacts_activateContactChannelCmd.MarkFlagRequired("activation-code")
		ssmContacts_activateContactChannelCmd.MarkFlagRequired("contact-channel-id")
	})
	ssmContactsCmd.AddCommand(ssmContacts_activateContactChannelCmd)
}
