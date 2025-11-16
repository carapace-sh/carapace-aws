package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_getUserPoolMfaConfigCmd = &cobra.Command{
	Use:   "get-user-pool-mfa-config",
	Short: "Given a user pool ID, returns configuration for sign-in with WebAuthn authenticators and for multi-factor authentication (MFA).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_getUserPoolMfaConfigCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoIdp_getUserPoolMfaConfigCmd).Standalone()

		cognitoIdp_getUserPoolMfaConfigCmd.Flags().String("user-pool-id", "", "The ID of the user pool where you want to query WebAuthn and MFA configuration.")
		cognitoIdp_getUserPoolMfaConfigCmd.MarkFlagRequired("user-pool-id")
	})
	cognitoIdpCmd.AddCommand(cognitoIdp_getUserPoolMfaConfigCmd)
}
