package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmContacts_deleteContactChannelCmd = &cobra.Command{
	Use:   "delete-contact-channel",
	Short: "To stop receiving engagements on a contact channel, you can delete the channel from a contact.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmContacts_deleteContactChannelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssmContacts_deleteContactChannelCmd).Standalone()

		ssmContacts_deleteContactChannelCmd.Flags().String("contact-channel-id", "", "The Amazon Resource Name (ARN) of the contact channel.")
		ssmContacts_deleteContactChannelCmd.MarkFlagRequired("contact-channel-id")
	})
	ssmContactsCmd.AddCommand(ssmContacts_deleteContactChannelCmd)
}
