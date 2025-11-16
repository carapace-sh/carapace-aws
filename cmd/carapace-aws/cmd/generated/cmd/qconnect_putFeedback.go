package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_putFeedbackCmd = &cobra.Command{
	Use:   "put-feedback",
	Short: "Provides feedback against the specified assistant for the specified target.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_putFeedbackCmd).Standalone()

	qconnect_putFeedbackCmd.Flags().String("assistant-id", "", "The identifier of the Amazon Q in Connect assistant.")
	qconnect_putFeedbackCmd.Flags().String("content-feedback", "", "Information about the feedback provided.")
	qconnect_putFeedbackCmd.Flags().String("target-id", "", "The identifier of the feedback target.")
	qconnect_putFeedbackCmd.Flags().String("target-type", "", "The type of the feedback target.")
	qconnect_putFeedbackCmd.MarkFlagRequired("assistant-id")
	qconnect_putFeedbackCmd.MarkFlagRequired("content-feedback")
	qconnect_putFeedbackCmd.MarkFlagRequired("target-id")
	qconnect_putFeedbackCmd.MarkFlagRequired("target-type")
	qconnectCmd.AddCommand(qconnect_putFeedbackCmd)
}
