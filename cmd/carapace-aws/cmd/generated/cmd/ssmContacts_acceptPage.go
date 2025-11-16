package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmContacts_acceptPageCmd = &cobra.Command{
	Use:   "accept-page",
	Short: "Used to acknowledge an engagement to a contact channel during an incident.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmContacts_acceptPageCmd).Standalone()

	ssmContacts_acceptPageCmd.Flags().String("accept-code", "", "A 6-digit code used to acknowledge the page.")
	ssmContacts_acceptPageCmd.Flags().String("accept-code-validation", "", "An optional field that Incident Manager uses to `ENFORCE` `AcceptCode` validation when acknowledging an page.")
	ssmContacts_acceptPageCmd.Flags().String("accept-type", "", "The type indicates if the page was `DELIVERED` or `READ`.")
	ssmContacts_acceptPageCmd.Flags().String("contact-channel-id", "", "The ARN of the contact channel.")
	ssmContacts_acceptPageCmd.Flags().String("note", "", "Information provided by the user when the user acknowledges the page.")
	ssmContacts_acceptPageCmd.Flags().String("page-id", "", "The Amazon Resource Name (ARN) of the engagement to a contact channel.")
	ssmContacts_acceptPageCmd.MarkFlagRequired("accept-code")
	ssmContacts_acceptPageCmd.MarkFlagRequired("accept-type")
	ssmContacts_acceptPageCmd.MarkFlagRequired("page-id")
	ssmContactsCmd.AddCommand(ssmContacts_acceptPageCmd)
}
