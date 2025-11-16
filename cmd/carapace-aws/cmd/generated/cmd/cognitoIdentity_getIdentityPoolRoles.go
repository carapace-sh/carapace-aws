package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdentity_getIdentityPoolRolesCmd = &cobra.Command{
	Use:   "get-identity-pool-roles",
	Short: "Gets the roles for an identity pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdentity_getIdentityPoolRolesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoIdentity_getIdentityPoolRolesCmd).Standalone()

		cognitoIdentity_getIdentityPoolRolesCmd.Flags().String("identity-pool-id", "", "An identity pool ID in the format REGION:GUID.")
		cognitoIdentity_getIdentityPoolRolesCmd.MarkFlagRequired("identity-pool-id")
	})
	cognitoIdentityCmd.AddCommand(cognitoIdentity_getIdentityPoolRolesCmd)
}
