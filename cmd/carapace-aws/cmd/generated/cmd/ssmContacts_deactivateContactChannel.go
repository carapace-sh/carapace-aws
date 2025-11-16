package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmContacts_deactivateContactChannelCmd = &cobra.Command{
	Use:   "deactivate-contact-channel",
	Short: "To no longer receive Incident Manager engagements to a contact channel, you can deactivate the channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmContacts_deactivateContactChannelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssmContacts_deactivateContactChannelCmd).Standalone()

		ssmContacts_deactivateContactChannelCmd.Flags().String("contact-channel-id", "", "The Amazon Resource Name (ARN) of the contact channel you're deactivating.")
		ssmContacts_deactivateContactChannelCmd.MarkFlagRequired("contact-channel-id")
	})
	ssmContactsCmd.AddCommand(ssmContacts_deactivateContactChannelCmd)
}
