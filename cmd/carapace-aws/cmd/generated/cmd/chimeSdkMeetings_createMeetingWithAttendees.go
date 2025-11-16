package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMeetings_createMeetingWithAttendeesCmd = &cobra.Command{
	Use:   "create-meeting-with-attendees",
	Short: "Creates a new Amazon Chime SDK meeting in the specified media Region, with attendees.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMeetings_createMeetingWithAttendeesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkMeetings_createMeetingWithAttendeesCmd).Standalone()

		chimeSdkMeetings_createMeetingWithAttendeesCmd.Flags().String("attendees", "", "The attendee information, including attendees' IDs and join tokens.")
		chimeSdkMeetings_createMeetingWithAttendeesCmd.Flags().String("client-request-token", "", "The unique identifier for the client request.")
		chimeSdkMeetings_createMeetingWithAttendeesCmd.Flags().String("external-meeting-id", "", "The external meeting ID.")
		chimeSdkMeetings_createMeetingWithAttendeesCmd.Flags().String("media-placement-network-type", "", "The type of network for the media placement.")
		chimeSdkMeetings_createMeetingWithAttendeesCmd.Flags().String("media-region", "", "The Region in which to create the meeting.")
		chimeSdkMeetings_createMeetingWithAttendeesCmd.Flags().String("meeting-features", "", "Lists the audio and video features enabled for a meeting, such as echo reduction.")
		chimeSdkMeetings_createMeetingWithAttendeesCmd.Flags().String("meeting-host-id", "", "Reserved.")
		chimeSdkMeetings_createMeetingWithAttendeesCmd.Flags().String("notifications-configuration", "", "The configuration for resource targets to receive notifications when meeting and attendee events occur.")
		chimeSdkMeetings_createMeetingWithAttendeesCmd.Flags().String("primary-meeting-id", "", "When specified, replicates the media from the primary meeting to the new meeting.")
		chimeSdkMeetings_createMeetingWithAttendeesCmd.Flags().String("tags", "", "The tags in the request.")
		chimeSdkMeetings_createMeetingWithAttendeesCmd.Flags().String("tenant-ids", "", "A consistent and opaque identifier, created and maintained by the builder to represent a segment of their users.")
		chimeSdkMeetings_createMeetingWithAttendeesCmd.MarkFlagRequired("attendees")
		chimeSdkMeetings_createMeetingWithAttendeesCmd.MarkFlagRequired("client-request-token")
		chimeSdkMeetings_createMeetingWithAttendeesCmd.MarkFlagRequired("external-meeting-id")
		chimeSdkMeetings_createMeetingWithAttendeesCmd.MarkFlagRequired("media-region")
	})
	chimeSdkMeetingsCmd.AddCommand(chimeSdkMeetings_createMeetingWithAttendeesCmd)
}
