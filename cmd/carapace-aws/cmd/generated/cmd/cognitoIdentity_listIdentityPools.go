package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdentity_listIdentityPoolsCmd = &cobra.Command{
	Use:   "list-identity-pools",
	Short: "Lists all of the Cognito identity pools registered for your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdentity_listIdentityPoolsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoIdentity_listIdentityPoolsCmd).Standalone()

		cognitoIdentity_listIdentityPoolsCmd.Flags().String("max-results", "", "The maximum number of identities to return.")
		cognitoIdentity_listIdentityPoolsCmd.Flags().String("next-token", "", "A pagination token.")
		cognitoIdentity_listIdentityPoolsCmd.MarkFlagRequired("max-results")
	})
	cognitoIdentityCmd.AddCommand(cognitoIdentity_listIdentityPoolsCmd)
}
