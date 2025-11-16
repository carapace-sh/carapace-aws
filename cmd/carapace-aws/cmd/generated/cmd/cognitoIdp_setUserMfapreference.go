package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_setUserMfapreferenceCmd = &cobra.Command{
	Use:   "set-user-mfapreference",
	Short: "Set the user's multi-factor authentication (MFA) method preference, including which MFA factors are activated and if any are preferred.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_setUserMfapreferenceCmd).Standalone()

	cognitoIdp_setUserMfapreferenceCmd.Flags().String("access-token", "", "A valid access token that Amazon Cognito issued to the currently signed-in user.")
	cognitoIdp_setUserMfapreferenceCmd.Flags().String("email-mfa-settings", "", "User preferences for email message MFA.")
	cognitoIdp_setUserMfapreferenceCmd.Flags().String("smsmfa-settings", "", "User preferences for SMS message MFA.")
	cognitoIdp_setUserMfapreferenceCmd.Flags().String("software-token-mfa-settings", "", "User preferences for time-based one-time password (TOTP) MFA.")
	cognitoIdp_setUserMfapreferenceCmd.MarkFlagRequired("access-token")
	cognitoIdpCmd.AddCommand(cognitoIdp_setUserMfapreferenceCmd)
}
