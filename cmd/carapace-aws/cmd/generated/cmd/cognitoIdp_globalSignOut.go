package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_globalSignOutCmd = &cobra.Command{
	Use:   "global-sign-out",
	Short: "Invalidates the identity, access, and refresh tokens that Amazon Cognito issued to a user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_globalSignOutCmd).Standalone()

	cognitoIdp_globalSignOutCmd.Flags().String("access-token", "", "A valid access token that Amazon Cognito issued to the currently signed-in user.")
	cognitoIdp_globalSignOutCmd.MarkFlagRequired("access-token")
	cognitoIdpCmd.AddCommand(cognitoIdp_globalSignOutCmd)
}
