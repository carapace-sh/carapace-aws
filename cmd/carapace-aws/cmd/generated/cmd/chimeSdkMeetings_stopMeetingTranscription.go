package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMeetings_stopMeetingTranscriptionCmd = &cobra.Command{
	Use:   "stop-meeting-transcription",
	Short: "Stops transcription for the specified `meetingId`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMeetings_stopMeetingTranscriptionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkMeetings_stopMeetingTranscriptionCmd).Standalone()

		chimeSdkMeetings_stopMeetingTranscriptionCmd.Flags().String("meeting-id", "", "The unique ID of the meeting for which you stop transcription.")
		chimeSdkMeetings_stopMeetingTranscriptionCmd.MarkFlagRequired("meeting-id")
	})
	chimeSdkMeetingsCmd.AddCommand(chimeSdkMeetings_stopMeetingTranscriptionCmd)
}
