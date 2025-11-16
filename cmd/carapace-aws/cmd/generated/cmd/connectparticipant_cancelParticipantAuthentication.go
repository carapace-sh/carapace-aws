package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectparticipant_cancelParticipantAuthenticationCmd = &cobra.Command{
	Use:   "cancel-participant-authentication",
	Short: "Cancels the authentication session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectparticipant_cancelParticipantAuthenticationCmd).Standalone()

	connectparticipant_cancelParticipantAuthenticationCmd.Flags().String("connection-token", "", "The authentication token associated with the participant's connection.")
	connectparticipant_cancelParticipantAuthenticationCmd.Flags().String("session-id", "", "The `sessionId` provided in the `authenticationInitiated` event.")
	connectparticipant_cancelParticipantAuthenticationCmd.MarkFlagRequired("connection-token")
	connectparticipant_cancelParticipantAuthenticationCmd.MarkFlagRequired("session-id")
	connectparticipantCmd.AddCommand(connectparticipant_cancelParticipantAuthenticationCmd)
}
