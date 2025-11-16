package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_updateParticipantAuthenticationCmd = &cobra.Command{
	Use:   "update-participant-authentication",
	Short: "Instructs Amazon Connect to resume the authentication process.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_updateParticipantAuthenticationCmd).Standalone()

	connect_updateParticipantAuthenticationCmd.Flags().String("code", "", "The `code` query parameter provided by Cognito in the `redirectUri`.")
	connect_updateParticipantAuthenticationCmd.Flags().String("error", "", "The `error` query parameter provided by Cognito in the `redirectUri`.")
	connect_updateParticipantAuthenticationCmd.Flags().String("error-description", "", "The `error_description` parameter provided by Cognito in the `redirectUri`.")
	connect_updateParticipantAuthenticationCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_updateParticipantAuthenticationCmd.Flags().String("state", "", "The `state` query parameter that was provided by Cognito in the `redirectUri`.")
	connect_updateParticipantAuthenticationCmd.MarkFlagRequired("instance-id")
	connect_updateParticipantAuthenticationCmd.MarkFlagRequired("state")
	connectCmd.AddCommand(connect_updateParticipantAuthenticationCmd)
}
