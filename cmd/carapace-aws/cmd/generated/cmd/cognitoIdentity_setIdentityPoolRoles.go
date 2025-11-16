package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdentity_setIdentityPoolRolesCmd = &cobra.Command{
	Use:   "set-identity-pool-roles",
	Short: "Sets the roles for an identity pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdentity_setIdentityPoolRolesCmd).Standalone()

	cognitoIdentity_setIdentityPoolRolesCmd.Flags().String("identity-pool-id", "", "An identity pool ID in the format REGION:GUID.")
	cognitoIdentity_setIdentityPoolRolesCmd.Flags().String("role-mappings", "", "How users for a specific identity provider are to mapped to roles.")
	cognitoIdentity_setIdentityPoolRolesCmd.Flags().String("roles", "", "The map of roles associated with this pool.")
	cognitoIdentity_setIdentityPoolRolesCmd.MarkFlagRequired("identity-pool-id")
	cognitoIdentity_setIdentityPoolRolesCmd.MarkFlagRequired("roles")
	cognitoIdentityCmd.AddCommand(cognitoIdentity_setIdentityPoolRolesCmd)
}
