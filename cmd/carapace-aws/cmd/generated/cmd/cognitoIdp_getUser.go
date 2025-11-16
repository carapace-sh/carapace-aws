package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_getUserCmd = &cobra.Command{
	Use:   "get-user",
	Short: "Gets user attributes and and MFA settings for the currently signed-in user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_getUserCmd).Standalone()

	cognitoIdp_getUserCmd.Flags().String("access-token", "", "A valid access token that Amazon Cognito issued to the currently signed-in user.")
	cognitoIdp_getUserCmd.MarkFlagRequired("access-token")
	cognitoIdpCmd.AddCommand(cognitoIdp_getUserCmd)
}
