package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdentity_getCredentialsForIdentityCmd = &cobra.Command{
	Use:   "get-credentials-for-identity",
	Short: "Returns credentials for the provided identity ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdentity_getCredentialsForIdentityCmd).Standalone()

	cognitoIdentity_getCredentialsForIdentityCmd.Flags().String("custom-role-arn", "", "The Amazon Resource Name (ARN) of the role to be assumed when multiple roles were received in the token from the identity provider.")
	cognitoIdentity_getCredentialsForIdentityCmd.Flags().String("identity-id", "", "A unique identifier in the format REGION:GUID.")
	cognitoIdentity_getCredentialsForIdentityCmd.Flags().String("logins", "", "A set of optional name-value pairs that map provider names to provider tokens.")
	cognitoIdentity_getCredentialsForIdentityCmd.MarkFlagRequired("identity-id")
	cognitoIdentityCmd.AddCommand(cognitoIdentity_getCredentialsForIdentityCmd)
}
