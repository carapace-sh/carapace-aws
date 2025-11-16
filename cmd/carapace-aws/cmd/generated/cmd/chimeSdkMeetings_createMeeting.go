package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMeetings_createMeetingCmd = &cobra.Command{
	Use:   "create-meeting",
	Short: "Creates a new Amazon Chime SDK meeting in the specified media Region with no initial attendees.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMeetings_createMeetingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkMeetings_createMeetingCmd).Standalone()

		chimeSdkMeetings_createMeetingCmd.Flags().String("client-request-token", "", "The unique identifier for the client request.")
		chimeSdkMeetings_createMeetingCmd.Flags().String("external-meeting-id", "", "The external meeting ID.")
		chimeSdkMeetings_createMeetingCmd.Flags().String("media-placement-network-type", "", "The type of network for the media placement.")
		chimeSdkMeetings_createMeetingCmd.Flags().String("media-region", "", "The Region in which to create the meeting.")
		chimeSdkMeetings_createMeetingCmd.Flags().String("meeting-features", "", "Lists the audio and video features enabled for a meeting, such as echo reduction.")
		chimeSdkMeetings_createMeetingCmd.Flags().String("meeting-host-id", "", "Reserved.")
		chimeSdkMeetings_createMeetingCmd.Flags().String("notifications-configuration", "", "The configuration for resource targets to receive notifications when meeting and attendee events occur.")
		chimeSdkMeetings_createMeetingCmd.Flags().String("primary-meeting-id", "", "When specified, replicates the media from the primary meeting to the new meeting.")
		chimeSdkMeetings_createMeetingCmd.Flags().String("tags", "", "Applies one or more tags to an Amazon Chime SDK meeting.")
		chimeSdkMeetings_createMeetingCmd.Flags().String("tenant-ids", "", "A consistent and opaque identifier, created and maintained by the builder to represent a segment of their users.")
		chimeSdkMeetings_createMeetingCmd.MarkFlagRequired("client-request-token")
		chimeSdkMeetings_createMeetingCmd.MarkFlagRequired("external-meeting-id")
		chimeSdkMeetings_createMeetingCmd.MarkFlagRequired("media-region")
	})
	chimeSdkMeetingsCmd.AddCommand(chimeSdkMeetings_createMeetingCmd)
}
