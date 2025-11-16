package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var support_addCommunicationToCaseCmd = &cobra.Command{
	Use:   "add-communication-to-case",
	Short: "Adds additional customer communication to an Amazon Web Services Support case.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(support_addCommunicationToCaseCmd).Standalone()

	support_addCommunicationToCaseCmd.Flags().String("attachment-set-id", "", "The ID of a set of one or more attachments for the communication to add to the case.")
	support_addCommunicationToCaseCmd.Flags().String("case-id", "", "The support case ID requested or returned in the call.")
	support_addCommunicationToCaseCmd.Flags().String("cc-email-addresses", "", "The email addresses in the CC line of an email to be added to the support case.")
	support_addCommunicationToCaseCmd.Flags().String("communication-body", "", "The body of an email communication to add to the support case.")
	support_addCommunicationToCaseCmd.MarkFlagRequired("communication-body")
	supportCmd.AddCommand(support_addCommunicationToCaseCmd)
}
