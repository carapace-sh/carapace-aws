package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMeetings_startMeetingTranscriptionCmd = &cobra.Command{
	Use:   "start-meeting-transcription",
	Short: "Starts transcription for the specified `meetingId`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMeetings_startMeetingTranscriptionCmd).Standalone()

	chimeSdkMeetings_startMeetingTranscriptionCmd.Flags().String("meeting-id", "", "The unique ID of the meeting being transcribed.")
	chimeSdkMeetings_startMeetingTranscriptionCmd.Flags().String("transcription-configuration", "", "The configuration for the current transcription operation.")
	chimeSdkMeetings_startMeetingTranscriptionCmd.MarkFlagRequired("meeting-id")
	chimeSdkMeetings_startMeetingTranscriptionCmd.MarkFlagRequired("transcription-configuration")
	chimeSdkMeetingsCmd.AddCommand(chimeSdkMeetings_startMeetingTranscriptionCmd)
}
