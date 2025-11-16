package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectparticipant_sendEventCmd = &cobra.Command{
	Use:   "send-event",
	Short: "The `application/vnd.amazonaws.connect.event.connection.acknowledged` ContentType is no longer maintained since December 31, 2024. This event has been migrated to the [CreateParticipantConnection](https://docs.aws.amazon.com/connect-participant/latest/APIReference/API_CreateParticipantConnection.html) API using the `ConnectParticipant` field.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectparticipant_sendEventCmd).Standalone()

	connectparticipant_sendEventCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	connectparticipant_sendEventCmd.Flags().String("connection-token", "", "The authentication token associated with the participant's connection.")
	connectparticipant_sendEventCmd.Flags().String("content", "", "The content of the event to be sent (for example, message text).")
	connectparticipant_sendEventCmd.Flags().String("content-type", "", "The content type of the request.")
	connectparticipant_sendEventCmd.MarkFlagRequired("connection-token")
	connectparticipant_sendEventCmd.MarkFlagRequired("content-type")
	connectparticipantCmd.AddCommand(connectparticipant_sendEventCmd)
}
