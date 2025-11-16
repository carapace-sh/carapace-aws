package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmContacts_updateContactCmd = &cobra.Command{
	Use:   "update-contact",
	Short: "Updates the contact or escalation plan specified.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmContacts_updateContactCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssmContacts_updateContactCmd).Standalone()

		ssmContacts_updateContactCmd.Flags().String("contact-id", "", "The Amazon Resource Name (ARN) of the contact or escalation plan you're updating.")
		ssmContacts_updateContactCmd.Flags().String("display-name", "", "The full name of the contact or escalation plan.")
		ssmContacts_updateContactCmd.Flags().String("plan", "", "A list of stages.")
		ssmContacts_updateContactCmd.MarkFlagRequired("contact-id")
	})
	ssmContactsCmd.AddCommand(ssmContacts_updateContactCmd)
}
