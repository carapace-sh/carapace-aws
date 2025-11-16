package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMeetings_updateAttendeeCapabilitiesCmd = &cobra.Command{
	Use:   "update-attendee-capabilities",
	Short: "The capabilities that you want to update.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMeetings_updateAttendeeCapabilitiesCmd).Standalone()

	chimeSdkMeetings_updateAttendeeCapabilitiesCmd.Flags().String("attendee-id", "", "The ID of the attendee associated with the update request.")
	chimeSdkMeetings_updateAttendeeCapabilitiesCmd.Flags().String("capabilities", "", "The capabilities that you want to update.")
	chimeSdkMeetings_updateAttendeeCapabilitiesCmd.Flags().String("meeting-id", "", "The ID of the meeting associated with the update request.")
	chimeSdkMeetings_updateAttendeeCapabilitiesCmd.MarkFlagRequired("attendee-id")
	chimeSdkMeetings_updateAttendeeCapabilitiesCmd.MarkFlagRequired("capabilities")
	chimeSdkMeetings_updateAttendeeCapabilitiesCmd.MarkFlagRequired("meeting-id")
	chimeSdkMeetingsCmd.AddCommand(chimeSdkMeetings_updateAttendeeCapabilitiesCmd)
}
