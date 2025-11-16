package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMeetings_batchCreateAttendeeCmd = &cobra.Command{
	Use:   "batch-create-attendee",
	Short: "Creates up to 100 attendees for an active Amazon Chime SDK meeting.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMeetings_batchCreateAttendeeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkMeetings_batchCreateAttendeeCmd).Standalone()

		chimeSdkMeetings_batchCreateAttendeeCmd.Flags().String("attendees", "", "The attendee information, including attendees' IDs and join tokens.")
		chimeSdkMeetings_batchCreateAttendeeCmd.Flags().String("meeting-id", "", "The Amazon Chime SDK ID of the meeting to which you're adding attendees.")
		chimeSdkMeetings_batchCreateAttendeeCmd.MarkFlagRequired("attendees")
		chimeSdkMeetings_batchCreateAttendeeCmd.MarkFlagRequired("meeting-id")
	})
	chimeSdkMeetingsCmd.AddCommand(chimeSdkMeetings_batchCreateAttendeeCmd)
}
