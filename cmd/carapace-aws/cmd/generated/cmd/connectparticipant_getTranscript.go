package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectparticipant_getTranscriptCmd = &cobra.Command{
	Use:   "get-transcript",
	Short: "Retrieves a transcript of the session, including details about any attachments.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectparticipant_getTranscriptCmd).Standalone()

	connectparticipant_getTranscriptCmd.Flags().String("connection-token", "", "The authentication token associated with the participant's connection.")
	connectparticipant_getTranscriptCmd.Flags().String("contact-id", "", "The contactId from the current contact chain for which transcript is needed.")
	connectparticipant_getTranscriptCmd.Flags().String("max-results", "", "The maximum number of results to return in the page.")
	connectparticipant_getTranscriptCmd.Flags().String("next-token", "", "The pagination token.")
	connectparticipant_getTranscriptCmd.Flags().String("scan-direction", "", "The direction from StartPosition from which to retrieve message.")
	connectparticipant_getTranscriptCmd.Flags().String("sort-order", "", "The sort order for the records.")
	connectparticipant_getTranscriptCmd.Flags().String("start-position", "", "A filtering option for where to start.")
	connectparticipant_getTranscriptCmd.MarkFlagRequired("connection-token")
	connectparticipantCmd.AddCommand(connectparticipant_getTranscriptCmd)
}
