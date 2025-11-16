package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_updateGroupCmd = &cobra.Command{
	Use:   "update-group",
	Short: "Given the name of a user pool group, updates any of the properties for precedence, IAM role, or description.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_updateGroupCmd).Standalone()

	cognitoIdp_updateGroupCmd.Flags().String("description", "", "A new description of the existing group.")
	cognitoIdp_updateGroupCmd.Flags().String("group-name", "", "The name of the group that you want to update.")
	cognitoIdp_updateGroupCmd.Flags().String("precedence", "", "A non-negative integer value that specifies the precedence of this group relative to the other groups that a user can belong to in the user pool.")
	cognitoIdp_updateGroupCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of an IAM role that you want to associate with the group.")
	cognitoIdp_updateGroupCmd.Flags().String("user-pool-id", "", "The ID of the user pool that contains the group you want to update.")
	cognitoIdp_updateGroupCmd.MarkFlagRequired("group-name")
	cognitoIdp_updateGroupCmd.MarkFlagRequired("user-pool-id")
	cognitoIdpCmd.AddCommand(cognitoIdp_updateGroupCmd)
}
