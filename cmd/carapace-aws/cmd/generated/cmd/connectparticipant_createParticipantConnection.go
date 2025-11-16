package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectparticipant_createParticipantConnectionCmd = &cobra.Command{
	Use:   "create-participant-connection",
	Short: "Creates the participant's connection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectparticipant_createParticipantConnectionCmd).Standalone()

	connectparticipant_createParticipantConnectionCmd.Flags().String("connect-participant", "", "Amazon Connect Participant is used to mark the participant as connected for customer participant in message streaming, as well as for agent or manager participant in non-streaming chats.")
	connectparticipant_createParticipantConnectionCmd.Flags().String("participant-token", "", "This is a header parameter.")
	connectparticipant_createParticipantConnectionCmd.Flags().String("type", "", "Type of connection information required.")
	connectparticipant_createParticipantConnectionCmd.MarkFlagRequired("participant-token")
	connectparticipantCmd.AddCommand(connectparticipant_createParticipantConnectionCmd)
}
