package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmContacts_getContactChannelCmd = &cobra.Command{
	Use:   "get-contact-channel",
	Short: "List details about a specific contact channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmContacts_getContactChannelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssmContacts_getContactChannelCmd).Standalone()

		ssmContacts_getContactChannelCmd.Flags().String("contact-channel-id", "", "The Amazon Resource Name (ARN) of the contact channel you want information about.")
		ssmContacts_getContactChannelCmd.MarkFlagRequired("contact-channel-id")
	})
	ssmContactsCmd.AddCommand(ssmContacts_getContactChannelCmd)
}
