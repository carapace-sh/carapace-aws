package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_listResourceServersCmd = &cobra.Command{
	Use:   "list-resource-servers",
	Short: "Given a user pool ID, returns all resource servers and their details.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_listResourceServersCmd).Standalone()

	cognitoIdp_listResourceServersCmd.Flags().String("max-results", "", "The maximum number of resource servers that you want Amazon Cognito to return in the response.")
	cognitoIdp_listResourceServersCmd.Flags().String("next-token", "", "This API operation returns a limited number of results.")
	cognitoIdp_listResourceServersCmd.Flags().String("user-pool-id", "", "The ID of the user pool where you want to list resource servers.")
	cognitoIdp_listResourceServersCmd.MarkFlagRequired("user-pool-id")
	cognitoIdpCmd.AddCommand(cognitoIdp_listResourceServersCmd)
}
