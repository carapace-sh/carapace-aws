package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmContacts_getContactCmd = &cobra.Command{
	Use:   "get-contact",
	Short: "Retrieves information about the specified contact or escalation plan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmContacts_getContactCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssmContacts_getContactCmd).Standalone()

		ssmContacts_getContactCmd.Flags().String("contact-id", "", "The Amazon Resource Name (ARN) of the contact or escalation plan.")
		ssmContacts_getContactCmd.MarkFlagRequired("contact-id")
	})
	ssmContactsCmd.AddCommand(ssmContacts_getContactCmd)
}
