package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmContacts_deleteContactCmd = &cobra.Command{
	Use:   "delete-contact",
	Short: "To remove a contact from Incident Manager, you can delete the contact.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmContacts_deleteContactCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssmContacts_deleteContactCmd).Standalone()

		ssmContacts_deleteContactCmd.Flags().String("contact-id", "", "The Amazon Resource Name (ARN) of the contact that you're deleting.")
		ssmContacts_deleteContactCmd.MarkFlagRequired("contact-id")
	})
	ssmContactsCmd.AddCommand(ssmContacts_deleteContactCmd)
}
