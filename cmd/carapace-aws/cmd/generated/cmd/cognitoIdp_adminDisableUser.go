package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_adminDisableUserCmd = &cobra.Command{
	Use:   "admin-disable-user",
	Short: "Deactivates a user profile and revokes all access tokens for the user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_adminDisableUserCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoIdp_adminDisableUserCmd).Standalone()

		cognitoIdp_adminDisableUserCmd.Flags().String("user-pool-id", "", "The ID of the user pool where you want to disable the user.")
		cognitoIdp_adminDisableUserCmd.Flags().String("username", "", "The name of the user that you want to query or modify.")
		cognitoIdp_adminDisableUserCmd.MarkFlagRequired("user-pool-id")
		cognitoIdp_adminDisableUserCmd.MarkFlagRequired("username")
	})
	cognitoIdpCmd.AddCommand(cognitoIdp_adminDisableUserCmd)
}
