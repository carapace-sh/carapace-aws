package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_setUserPoolMfaConfigCmd = &cobra.Command{
	Use:   "set-user-pool-mfa-config",
	Short: "Sets user pool multi-factor authentication (MFA) and passkey configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_setUserPoolMfaConfigCmd).Standalone()

	cognitoIdp_setUserPoolMfaConfigCmd.Flags().String("email-mfa-configuration", "", "Sets configuration for user pool email message MFA and sign-in with one-time passwords (OTPs).")
	cognitoIdp_setUserPoolMfaConfigCmd.Flags().String("mfa-configuration", "", "Sets multi-factor authentication (MFA) to be on, off, or optional.")
	cognitoIdp_setUserPoolMfaConfigCmd.Flags().String("sms-mfa-configuration", "", "Configures user pool SMS messages for MFA.")
	cognitoIdp_setUserPoolMfaConfigCmd.Flags().String("software-token-mfa-configuration", "", "Configures a user pool for time-based one-time password (TOTP) MFA.")
	cognitoIdp_setUserPoolMfaConfigCmd.Flags().String("user-pool-id", "", "The user pool ID.")
	cognitoIdp_setUserPoolMfaConfigCmd.Flags().String("web-authn-configuration", "", "The configuration of your user pool for passkey, or WebAuthn, authentication and registration.")
	cognitoIdp_setUserPoolMfaConfigCmd.MarkFlagRequired("user-pool-id")
	cognitoIdpCmd.AddCommand(cognitoIdp_setUserPoolMfaConfigCmd)
}
