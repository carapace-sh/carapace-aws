package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectparticipant_getAuthenticationUrlCmd = &cobra.Command{
	Use:   "get-authentication-url",
	Short: "Retrieves the AuthenticationUrl for the current authentication session for the AuthenticateCustomer flow block.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectparticipant_getAuthenticationUrlCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connectparticipant_getAuthenticationUrlCmd).Standalone()

		connectparticipant_getAuthenticationUrlCmd.Flags().String("connection-token", "", "The authentication token associated with the participant's connection.")
		connectparticipant_getAuthenticationUrlCmd.Flags().String("redirect-uri", "", "The URL where the customer will be redirected after Amazon Cognito authorizes the user.")
		connectparticipant_getAuthenticationUrlCmd.Flags().String("session-id", "", "The sessionId provided in the authenticationInitiated event.")
		connectparticipant_getAuthenticationUrlCmd.MarkFlagRequired("connection-token")
		connectparticipant_getAuthenticationUrlCmd.MarkFlagRequired("redirect-uri")
		connectparticipant_getAuthenticationUrlCmd.MarkFlagRequired("session-id")
	})
	connectparticipantCmd.AddCommand(connectparticipant_getAuthenticationUrlCmd)
}
