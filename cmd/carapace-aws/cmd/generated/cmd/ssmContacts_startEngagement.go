package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmContacts_startEngagementCmd = &cobra.Command{
	Use:   "start-engagement",
	Short: "Starts an engagement to a contact or escalation plan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmContacts_startEngagementCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssmContacts_startEngagementCmd).Standalone()

		ssmContacts_startEngagementCmd.Flags().String("contact-id", "", "The Amazon Resource Name (ARN) of the contact being engaged.")
		ssmContacts_startEngagementCmd.Flags().String("content", "", "The secure content of the message that was sent to the contact.")
		ssmContacts_startEngagementCmd.Flags().String("idempotency-token", "", "A token ensuring that the operation is called only once with the specified details.")
		ssmContacts_startEngagementCmd.Flags().String("incident-id", "", "The ARN of the incident that the engagement is part of.")
		ssmContacts_startEngagementCmd.Flags().String("public-content", "", "The insecure content of the message that was sent to the contact.")
		ssmContacts_startEngagementCmd.Flags().String("public-subject", "", "The insecure subject of the message that was sent to the contact.")
		ssmContacts_startEngagementCmd.Flags().String("sender", "", "The user that started the engagement.")
		ssmContacts_startEngagementCmd.Flags().String("subject", "", "The secure subject of the message that was sent to the contact.")
		ssmContacts_startEngagementCmd.MarkFlagRequired("contact-id")
		ssmContacts_startEngagementCmd.MarkFlagRequired("content")
		ssmContacts_startEngagementCmd.MarkFlagRequired("sender")
		ssmContacts_startEngagementCmd.MarkFlagRequired("subject")
	})
	ssmContactsCmd.AddCommand(ssmContacts_startEngagementCmd)
}
