package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chime_createMeetingDialOutCmd = &cobra.Command{
	Use:   "create-meeting-dial-out",
	Short: "Uses the join token and call metadata in a meeting request (From number, To number, and so forth) to initiate an outbound call to a public switched telephone network (PSTN) and join them into a Chime meeting.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chime_createMeetingDialOutCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chime_createMeetingDialOutCmd).Standalone()

		chime_createMeetingDialOutCmd.Flags().String("from-phone-number", "", "Phone number used as the caller ID when the remote party receives a call.")
		chime_createMeetingDialOutCmd.Flags().String("join-token", "", "Token used by the Amazon Chime SDK attendee.")
		chime_createMeetingDialOutCmd.Flags().String("meeting-id", "", "The Amazon Chime SDK meeting ID.")
		chime_createMeetingDialOutCmd.Flags().String("to-phone-number", "", "Phone number called when inviting someone to a meeting.")
		chime_createMeetingDialOutCmd.MarkFlagRequired("from-phone-number")
		chime_createMeetingDialOutCmd.MarkFlagRequired("join-token")
		chime_createMeetingDialOutCmd.MarkFlagRequired("meeting-id")
		chime_createMeetingDialOutCmd.MarkFlagRequired("to-phone-number")
	})
	chimeCmd.AddCommand(chime_createMeetingDialOutCmd)
}
