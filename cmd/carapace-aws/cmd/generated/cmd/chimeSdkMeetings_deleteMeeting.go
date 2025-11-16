package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMeetings_deleteMeetingCmd = &cobra.Command{
	Use:   "delete-meeting",
	Short: "Deletes the specified Amazon Chime SDK meeting.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMeetings_deleteMeetingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkMeetings_deleteMeetingCmd).Standalone()

		chimeSdkMeetings_deleteMeetingCmd.Flags().String("meeting-id", "", "The Amazon Chime SDK meeting ID.")
		chimeSdkMeetings_deleteMeetingCmd.MarkFlagRequired("meeting-id")
	})
	chimeSdkMeetingsCmd.AddCommand(chimeSdkMeetings_deleteMeetingCmd)
}
