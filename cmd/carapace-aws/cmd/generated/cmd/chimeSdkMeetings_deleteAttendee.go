package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMeetings_deleteAttendeeCmd = &cobra.Command{
	Use:   "delete-attendee",
	Short: "Deletes an attendee from the specified Amazon Chime SDK meeting and deletes their `JoinToken`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMeetings_deleteAttendeeCmd).Standalone()

	chimeSdkMeetings_deleteAttendeeCmd.Flags().String("attendee-id", "", "The Amazon Chime SDK attendee ID.")
	chimeSdkMeetings_deleteAttendeeCmd.Flags().String("meeting-id", "", "The Amazon Chime SDK meeting ID.")
	chimeSdkMeetings_deleteAttendeeCmd.MarkFlagRequired("attendee-id")
	chimeSdkMeetings_deleteAttendeeCmd.MarkFlagRequired("meeting-id")
	chimeSdkMeetingsCmd.AddCommand(chimeSdkMeetings_deleteAttendeeCmd)
}
