package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_deleteUserCmd = &cobra.Command{
	Use:   "delete-user",
	Short: "Deletes the profile of the currently signed-in user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_deleteUserCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoIdp_deleteUserCmd).Standalone()

		cognitoIdp_deleteUserCmd.Flags().String("access-token", "", "A valid access token that Amazon Cognito issued to the currently signed-in user.")
		cognitoIdp_deleteUserCmd.MarkFlagRequired("access-token")
	})
	cognitoIdpCmd.AddCommand(cognitoIdp_deleteUserCmd)
}
