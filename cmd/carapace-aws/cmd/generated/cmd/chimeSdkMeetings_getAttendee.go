package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMeetings_getAttendeeCmd = &cobra.Command{
	Use:   "get-attendee",
	Short: "Gets the Amazon Chime SDK attendee details for a specified meeting ID and attendee ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMeetings_getAttendeeCmd).Standalone()

	chimeSdkMeetings_getAttendeeCmd.Flags().String("attendee-id", "", "The Amazon Chime SDK attendee ID.")
	chimeSdkMeetings_getAttendeeCmd.Flags().String("meeting-id", "", "The Amazon Chime SDK meeting ID.")
	chimeSdkMeetings_getAttendeeCmd.MarkFlagRequired("attendee-id")
	chimeSdkMeetings_getAttendeeCmd.MarkFlagRequired("meeting-id")
	chimeSdkMeetingsCmd.AddCommand(chimeSdkMeetings_getAttendeeCmd)
}
