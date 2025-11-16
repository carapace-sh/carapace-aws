package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmContacts_createContactCmd = &cobra.Command{
	Use:   "create-contact",
	Short: "Contacts are either the contacts that Incident Manager engages during an incident or the escalation plans that Incident Manager uses to engage contacts in phases during an incident.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmContacts_createContactCmd).Standalone()

	ssmContacts_createContactCmd.Flags().String("alias", "", "The short name to quickly identify a contact or escalation plan.")
	ssmContacts_createContactCmd.Flags().String("display-name", "", "The full name of the contact or escalation plan.")
	ssmContacts_createContactCmd.Flags().String("idempotency-token", "", "A token ensuring that the operation is called only once with the specified details.")
	ssmContacts_createContactCmd.Flags().String("plan", "", "A list of stages.")
	ssmContacts_createContactCmd.Flags().String("tags", "", "Adds a tag to the target.")
	ssmContacts_createContactCmd.Flags().String("type", "", "The type of contact to create.")
	ssmContacts_createContactCmd.MarkFlagRequired("alias")
	ssmContacts_createContactCmd.MarkFlagRequired("plan")
	ssmContacts_createContactCmd.MarkFlagRequired("type")
	ssmContactsCmd.AddCommand(ssmContacts_createContactCmd)
}
