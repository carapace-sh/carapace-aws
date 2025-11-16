package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmContacts_putContactPolicyCmd = &cobra.Command{
	Use:   "put-contact-policy",
	Short: "Adds a resource policy to the specified contact or escalation plan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmContacts_putContactPolicyCmd).Standalone()

	ssmContacts_putContactPolicyCmd.Flags().String("contact-arn", "", "The Amazon Resource Name (ARN) of the contact or escalation plan.")
	ssmContacts_putContactPolicyCmd.Flags().String("policy", "", "Details of the resource policy.")
	ssmContacts_putContactPolicyCmd.MarkFlagRequired("contact-arn")
	ssmContacts_putContactPolicyCmd.MarkFlagRequired("policy")
	ssmContactsCmd.AddCommand(ssmContacts_putContactPolicyCmd)
}
