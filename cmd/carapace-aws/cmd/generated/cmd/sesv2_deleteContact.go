package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_deleteContactCmd = &cobra.Command{
	Use:   "delete-contact",
	Short: "Removes a contact from a contact list.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_deleteContactCmd).Standalone()

	sesv2_deleteContactCmd.Flags().String("contact-list-name", "", "The name of the contact list from which the contact should be removed.")
	sesv2_deleteContactCmd.Flags().String("email-address", "", "The contact's email address.")
	sesv2_deleteContactCmd.MarkFlagRequired("contact-list-name")
	sesv2_deleteContactCmd.MarkFlagRequired("email-address")
	sesv2Cmd.AddCommand(sesv2_deleteContactCmd)
}
