package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_getUserAuthFactorsCmd = &cobra.Command{
	Use:   "get-user-auth-factors",
	Short: "Lists the authentication options for the currently signed-in user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_getUserAuthFactorsCmd).Standalone()

	cognitoIdp_getUserAuthFactorsCmd.Flags().String("access-token", "", "A valid access token that Amazon Cognito issued to the currently signed-in user.")
	cognitoIdp_getUserAuthFactorsCmd.MarkFlagRequired("access-token")
	cognitoIdpCmd.AddCommand(cognitoIdp_getUserAuthFactorsCmd)
}
