package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_deleteGroupCmd = &cobra.Command{
	Use:   "delete-group",
	Short: "Deletes a group from the specified user pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_deleteGroupCmd).Standalone()

	cognitoIdp_deleteGroupCmd.Flags().String("group-name", "", "The name of the group that you want to delete.")
	cognitoIdp_deleteGroupCmd.Flags().String("user-pool-id", "", "The ID of the user pool where you want to delete the group.")
	cognitoIdp_deleteGroupCmd.MarkFlagRequired("group-name")
	cognitoIdp_deleteGroupCmd.MarkFlagRequired("user-pool-id")
	cognitoIdpCmd.AddCommand(cognitoIdp_deleteGroupCmd)
}
