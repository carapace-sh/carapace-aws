package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectparticipant_disconnectParticipantCmd = &cobra.Command{
	Use:   "disconnect-participant",
	Short: "Disconnects a participant.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectparticipant_disconnectParticipantCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connectparticipant_disconnectParticipantCmd).Standalone()

		connectparticipant_disconnectParticipantCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		connectparticipant_disconnectParticipantCmd.Flags().String("connection-token", "", "The authentication token associated with the participant's connection.")
		connectparticipant_disconnectParticipantCmd.MarkFlagRequired("connection-token")
	})
	connectparticipantCmd.AddCommand(connectparticipant_disconnectParticipantCmd)
}
