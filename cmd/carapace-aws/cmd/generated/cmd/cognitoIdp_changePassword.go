package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_changePasswordCmd = &cobra.Command{
	Use:   "change-password",
	Short: "Changes the password for the currently signed-in user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_changePasswordCmd).Standalone()

	cognitoIdp_changePasswordCmd.Flags().String("access-token", "", "A valid access token that Amazon Cognito issued to the user whose password you want to change.")
	cognitoIdp_changePasswordCmd.Flags().String("previous-password", "", "The user's previous password.")
	cognitoIdp_changePasswordCmd.Flags().String("proposed-password", "", "A new password that you prompted the user to enter in your application.")
	cognitoIdp_changePasswordCmd.MarkFlagRequired("access-token")
	cognitoIdp_changePasswordCmd.MarkFlagRequired("proposed-password")
	cognitoIdpCmd.AddCommand(cognitoIdp_changePasswordCmd)
}
