package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_listIdentityProvidersCmd = &cobra.Command{
	Use:   "list-identity-providers",
	Short: "Given a user pool ID, returns information about configured identity providers (IdPs).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_listIdentityProvidersCmd).Standalone()

	cognitoIdp_listIdentityProvidersCmd.Flags().String("max-results", "", "The maximum number of IdPs that you want Amazon Cognito to return in the response.")
	cognitoIdp_listIdentityProvidersCmd.Flags().String("next-token", "", "This API operation returns a limited number of results.")
	cognitoIdp_listIdentityProvidersCmd.Flags().String("user-pool-id", "", "The ID of the user pool where you want to list IdPs.")
	cognitoIdp_listIdentityProvidersCmd.MarkFlagRequired("user-pool-id")
	cognitoIdpCmd.AddCommand(cognitoIdp_listIdentityProvidersCmd)
}
