package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_getMessageInsightsCmd = &cobra.Command{
	Use:   "get-message-insights",
	Short: "Provides information about a specific message, including the from address, the subject, the recipient address, email tags, as well as events associated with the message.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_getMessageInsightsCmd).Standalone()

	sesv2_getMessageInsightsCmd.Flags().String("message-id", "", "A `MessageId` is a unique identifier for a message, and is returned when sending emails through Amazon SES.")
	sesv2_getMessageInsightsCmd.MarkFlagRequired("message-id")
	sesv2Cmd.AddCommand(sesv2_getMessageInsightsCmd)
}
