package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_adminSetUserMfapreferenceCmd = &cobra.Command{
	Use:   "admin-set-user-mfapreference",
	Short: "Sets the user's multi-factor authentication (MFA) preference, including which MFA options are activated, and if any are preferred.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_adminSetUserMfapreferenceCmd).Standalone()

	cognitoIdp_adminSetUserMfapreferenceCmd.Flags().String("email-mfa-settings", "", "User preferences for email message MFA.")
	cognitoIdp_adminSetUserMfapreferenceCmd.Flags().String("smsmfa-settings", "", "User preferences for SMS message MFA.")
	cognitoIdp_adminSetUserMfapreferenceCmd.Flags().String("software-token-mfa-settings", "", "User preferences for time-based one-time password (TOTP) MFA.")
	cognitoIdp_adminSetUserMfapreferenceCmd.Flags().String("user-pool-id", "", "The ID of the user pool where you want to set a user's MFA preferences.")
	cognitoIdp_adminSetUserMfapreferenceCmd.Flags().String("username", "", "The name of the user that you want to query or modify.")
	cognitoIdp_adminSetUserMfapreferenceCmd.MarkFlagRequired("user-pool-id")
	cognitoIdp_adminSetUserMfapreferenceCmd.MarkFlagRequired("username")
	cognitoIdpCmd.AddCommand(cognitoIdp_adminSetUserMfapreferenceCmd)
}
