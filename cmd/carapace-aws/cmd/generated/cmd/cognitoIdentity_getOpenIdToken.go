package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdentity_getOpenIdTokenCmd = &cobra.Command{
	Use:   "get-open-id-token",
	Short: "Gets an OpenID token, using a known Cognito ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdentity_getOpenIdTokenCmd).Standalone()

	cognitoIdentity_getOpenIdTokenCmd.Flags().String("identity-id", "", "A unique identifier in the format REGION:GUID.")
	cognitoIdentity_getOpenIdTokenCmd.Flags().String("logins", "", "A set of optional name-value pairs that map provider names to provider tokens.")
	cognitoIdentity_getOpenIdTokenCmd.MarkFlagRequired("identity-id")
	cognitoIdentityCmd.AddCommand(cognitoIdentity_getOpenIdTokenCmd)
}
