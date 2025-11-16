package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmContacts_getContactPolicyCmd = &cobra.Command{
	Use:   "get-contact-policy",
	Short: "Retrieves the resource policies attached to the specified contact or escalation plan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmContacts_getContactPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssmContacts_getContactPolicyCmd).Standalone()

		ssmContacts_getContactPolicyCmd.Flags().String("contact-arn", "", "The Amazon Resource Name (ARN) of the contact or escalation plan.")
		ssmContacts_getContactPolicyCmd.MarkFlagRequired("contact-arn")
	})
	ssmContactsCmd.AddCommand(ssmContacts_getContactPolicyCmd)
}
