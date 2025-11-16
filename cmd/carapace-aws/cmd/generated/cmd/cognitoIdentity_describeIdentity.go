package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdentity_describeIdentityCmd = &cobra.Command{
	Use:   "describe-identity",
	Short: "Returns metadata related to the given identity, including when the identity was created and any associated linked logins.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdentity_describeIdentityCmd).Standalone()

	cognitoIdentity_describeIdentityCmd.Flags().String("identity-id", "", "A unique identifier in the format REGION:GUID.")
	cognitoIdentity_describeIdentityCmd.MarkFlagRequired("identity-id")
	cognitoIdentityCmd.AddCommand(cognitoIdentity_describeIdentityCmd)
}
