package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmContacts_sendActivationCodeCmd = &cobra.Command{
	Use:   "send-activation-code",
	Short: "Sends an activation code to a contact channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmContacts_sendActivationCodeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssmContacts_sendActivationCodeCmd).Standalone()

		ssmContacts_sendActivationCodeCmd.Flags().String("contact-channel-id", "", "The Amazon Resource Name (ARN) of the contact channel.")
		ssmContacts_sendActivationCodeCmd.MarkFlagRequired("contact-channel-id")
	})
	ssmContactsCmd.AddCommand(ssmContacts_sendActivationCodeCmd)
}
