package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMeetings_batchUpdateAttendeeCapabilitiesExceptCmd = &cobra.Command{
	Use:   "batch-update-attendee-capabilities-except",
	Short: "Updates `AttendeeCapabilities` except the capabilities listed in an `ExcludedAttendeeIds` table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMeetings_batchUpdateAttendeeCapabilitiesExceptCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkMeetings_batchUpdateAttendeeCapabilitiesExceptCmd).Standalone()

		chimeSdkMeetings_batchUpdateAttendeeCapabilitiesExceptCmd.Flags().String("capabilities", "", "The capabilities (`audio`, `video`, or `content`) that you want to update.")
		chimeSdkMeetings_batchUpdateAttendeeCapabilitiesExceptCmd.Flags().String("excluded-attendee-ids", "", "The `AttendeeIDs` that you want to exclude from one or more capabilities.")
		chimeSdkMeetings_batchUpdateAttendeeCapabilitiesExceptCmd.Flags().String("meeting-id", "", "The ID of the meeting associated with the update request.")
		chimeSdkMeetings_batchUpdateAttendeeCapabilitiesExceptCmd.MarkFlagRequired("capabilities")
		chimeSdkMeetings_batchUpdateAttendeeCapabilitiesExceptCmd.MarkFlagRequired("excluded-attendee-ids")
		chimeSdkMeetings_batchUpdateAttendeeCapabilitiesExceptCmd.MarkFlagRequired("meeting-id")
	})
	chimeSdkMeetingsCmd.AddCommand(chimeSdkMeetings_batchUpdateAttendeeCapabilitiesExceptCmd)
}
