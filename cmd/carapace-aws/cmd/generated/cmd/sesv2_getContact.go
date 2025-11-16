package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_getContactCmd = &cobra.Command{
	Use:   "get-contact",
	Short: "Returns a contact from a contact list.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_getContactCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sesv2_getContactCmd).Standalone()

		sesv2_getContactCmd.Flags().String("contact-list-name", "", "The name of the contact list to which the contact belongs.")
		sesv2_getContactCmd.Flags().String("email-address", "", "The contact's email address.")
		sesv2_getContactCmd.MarkFlagRequired("contact-list-name")
		sesv2_getContactCmd.MarkFlagRequired("email-address")
	})
	sesv2Cmd.AddCommand(sesv2_getContactCmd)
}
