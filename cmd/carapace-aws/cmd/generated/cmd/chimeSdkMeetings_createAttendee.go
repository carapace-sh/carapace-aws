package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMeetings_createAttendeeCmd = &cobra.Command{
	Use:   "create-attendee",
	Short: "Creates a new attendee for an active Amazon Chime SDK meeting.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMeetings_createAttendeeCmd).Standalone()

	chimeSdkMeetings_createAttendeeCmd.Flags().String("capabilities", "", "The capabilities (`audio`, `video`, or `content`) that you want to grant an attendee.")
	chimeSdkMeetings_createAttendeeCmd.Flags().String("external-user-id", "", "The Amazon Chime SDK external user ID.")
	chimeSdkMeetings_createAttendeeCmd.Flags().String("meeting-id", "", "The unique ID of the meeting.")
	chimeSdkMeetings_createAttendeeCmd.MarkFlagRequired("external-user-id")
	chimeSdkMeetings_createAttendeeCmd.MarkFlagRequired("meeting-id")
	chimeSdkMeetingsCmd.AddCommand(chimeSdkMeetings_createAttendeeCmd)
}
