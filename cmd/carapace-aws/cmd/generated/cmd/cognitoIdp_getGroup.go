package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_getGroupCmd = &cobra.Command{
	Use:   "get-group",
	Short: "Given a user pool ID and a group name, returns information about the user group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_getGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoIdp_getGroupCmd).Standalone()

		cognitoIdp_getGroupCmd.Flags().String("group-name", "", "The name of the group that you want to get information about.")
		cognitoIdp_getGroupCmd.Flags().String("user-pool-id", "", "The ID of the user pool that contains the group that you want to query.")
		cognitoIdp_getGroupCmd.MarkFlagRequired("group-name")
		cognitoIdp_getGroupCmd.MarkFlagRequired("user-pool-id")
	})
	cognitoIdpCmd.AddCommand(cognitoIdp_getGroupCmd)
}
