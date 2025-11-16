package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_createParticipantCmd = &cobra.Command{
	Use:   "create-participant",
	Short: "Adds a new participant into an on-going chat contact or webRTC call.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_createParticipantCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_createParticipantCmd).Standalone()

		connect_createParticipantCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		connect_createParticipantCmd.Flags().String("contact-id", "", "The identifier of the contact in this instance of Amazon Connect.")
		connect_createParticipantCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_createParticipantCmd.Flags().String("participant-details", "", "Information identifying the participant.")
		connect_createParticipantCmd.MarkFlagRequired("contact-id")
		connect_createParticipantCmd.MarkFlagRequired("instance-id")
		connect_createParticipantCmd.MarkFlagRequired("participant-details")
	})
	connectCmd.AddCommand(connect_createParticipantCmd)
}
