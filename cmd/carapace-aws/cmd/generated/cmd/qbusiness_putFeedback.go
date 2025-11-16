package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_putFeedbackCmd = &cobra.Command{
	Use:   "put-feedback",
	Short: "Enables your end user to provide feedback on their Amazon Q Business generated chat responses.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_putFeedbackCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qbusiness_putFeedbackCmd).Standalone()

		qbusiness_putFeedbackCmd.Flags().String("application-id", "", "The identifier of the application associated with the feedback.")
		qbusiness_putFeedbackCmd.Flags().String("conversation-id", "", "The identifier of the conversation the feedback is attached to.")
		qbusiness_putFeedbackCmd.Flags().String("message-copied-at", "", "The timestamp for when the feedback was recorded.")
		qbusiness_putFeedbackCmd.Flags().String("message-id", "", "The identifier of the chat message that the feedback was given for.")
		qbusiness_putFeedbackCmd.Flags().String("message-usefulness", "", "The feedback usefulness value given by the user to the chat message.")
		qbusiness_putFeedbackCmd.Flags().String("user-id", "", "The identifier of the user giving the feedback.")
		qbusiness_putFeedbackCmd.MarkFlagRequired("application-id")
		qbusiness_putFeedbackCmd.MarkFlagRequired("conversation-id")
		qbusiness_putFeedbackCmd.MarkFlagRequired("message-id")
	})
	qbusinessCmd.AddCommand(qbusiness_putFeedbackCmd)
}
