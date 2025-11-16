package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_setUserSettingsCmd = &cobra.Command{
	Use:   "set-user-settings",
	Short: "*This action is no longer supported.* You can use it to configure only SMS MFA.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_setUserSettingsCmd).Standalone()

	cognitoIdp_setUserSettingsCmd.Flags().String("access-token", "", "A valid access token that Amazon Cognito issued to the currently signed-in user.")
	cognitoIdp_setUserSettingsCmd.Flags().String("mfaoptions", "", "You can use this parameter only to set an SMS configuration that uses SMS for delivery.")
	cognitoIdp_setUserSettingsCmd.MarkFlagRequired("access-token")
	cognitoIdp_setUserSettingsCmd.MarkFlagRequired("mfaoptions")
	cognitoIdpCmd.AddCommand(cognitoIdp_setUserSettingsCmd)
}
