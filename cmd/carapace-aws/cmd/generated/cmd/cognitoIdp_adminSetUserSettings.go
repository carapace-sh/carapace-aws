package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_adminSetUserSettingsCmd = &cobra.Command{
	Use:   "admin-set-user-settings",
	Short: "*This action is no longer supported.* You can use it to configure only SMS MFA.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_adminSetUserSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoIdp_adminSetUserSettingsCmd).Standalone()

		cognitoIdp_adminSetUserSettingsCmd.Flags().String("mfaoptions", "", "You can use this parameter only to set an SMS configuration that uses SMS for delivery.")
		cognitoIdp_adminSetUserSettingsCmd.Flags().String("user-pool-id", "", "The ID of the user pool that contains the user whose options you're setting.")
		cognitoIdp_adminSetUserSettingsCmd.Flags().String("username", "", "The name of the user that you want to query or modify.")
		cognitoIdp_adminSetUserSettingsCmd.MarkFlagRequired("mfaoptions")
		cognitoIdp_adminSetUserSettingsCmd.MarkFlagRequired("user-pool-id")
		cognitoIdp_adminSetUserSettingsCmd.MarkFlagRequired("username")
	})
	cognitoIdpCmd.AddCommand(cognitoIdp_adminSetUserSettingsCmd)
}
