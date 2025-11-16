package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdentity_unlinkIdentityCmd = &cobra.Command{
	Use:   "unlink-identity",
	Short: "Unlinks a federated identity from an existing account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdentity_unlinkIdentityCmd).Standalone()

	cognitoIdentity_unlinkIdentityCmd.Flags().String("identity-id", "", "A unique identifier in the format REGION:GUID.")
	cognitoIdentity_unlinkIdentityCmd.Flags().String("logins", "", "A set of optional name-value pairs that map provider names to provider tokens.")
	cognitoIdentity_unlinkIdentityCmd.Flags().String("logins-to-remove", "", "Provider names to unlink from this identity.")
	cognitoIdentity_unlinkIdentityCmd.MarkFlagRequired("identity-id")
	cognitoIdentity_unlinkIdentityCmd.MarkFlagRequired("logins")
	cognitoIdentity_unlinkIdentityCmd.MarkFlagRequired("logins-to-remove")
	cognitoIdentityCmd.AddCommand(cognitoIdentity_unlinkIdentityCmd)
}
