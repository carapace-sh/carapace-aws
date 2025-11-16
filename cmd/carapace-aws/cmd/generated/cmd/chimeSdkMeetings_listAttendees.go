package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMeetings_listAttendeesCmd = &cobra.Command{
	Use:   "list-attendees",
	Short: "Lists the attendees for the specified Amazon Chime SDK meeting.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMeetings_listAttendeesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkMeetings_listAttendeesCmd).Standalone()

		chimeSdkMeetings_listAttendeesCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
		chimeSdkMeetings_listAttendeesCmd.Flags().String("meeting-id", "", "The Amazon Chime SDK meeting ID.")
		chimeSdkMeetings_listAttendeesCmd.Flags().String("next-token", "", "The token to use to retrieve the next page of results.")
		chimeSdkMeetings_listAttendeesCmd.MarkFlagRequired("meeting-id")
	})
	chimeSdkMeetingsCmd.AddCommand(chimeSdkMeetings_listAttendeesCmd)
}
