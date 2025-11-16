package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMeetings_getMeetingCmd = &cobra.Command{
	Use:   "get-meeting",
	Short: "Gets the Amazon Chime SDK meeting details for the specified meeting ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMeetings_getMeetingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkMeetings_getMeetingCmd).Standalone()

		chimeSdkMeetings_getMeetingCmd.Flags().String("meeting-id", "", "The Amazon Chime SDK meeting ID.")
		chimeSdkMeetings_getMeetingCmd.MarkFlagRequired("meeting-id")
	})
	chimeSdkMeetingsCmd.AddCommand(chimeSdkMeetings_getMeetingCmd)
}
