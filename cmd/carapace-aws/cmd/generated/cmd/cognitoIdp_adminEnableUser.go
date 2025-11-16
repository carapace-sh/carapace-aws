package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_adminEnableUserCmd = &cobra.Command{
	Use:   "admin-enable-user",
	Short: "Activates sign-in for a user profile that previously had sign-in access disabled.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_adminEnableUserCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoIdp_adminEnableUserCmd).Standalone()

		cognitoIdp_adminEnableUserCmd.Flags().String("user-pool-id", "", "The ID of the user pool where you want to activate sign-in for the user.")
		cognitoIdp_adminEnableUserCmd.Flags().String("username", "", "The name of the user that you want to query or modify.")
		cognitoIdp_adminEnableUserCmd.MarkFlagRequired("user-pool-id")
		cognitoIdp_adminEnableUserCmd.MarkFlagRequired("username")
	})
	cognitoIdpCmd.AddCommand(cognitoIdp_adminEnableUserCmd)
}
