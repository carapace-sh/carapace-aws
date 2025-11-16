package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_createGroupCmd = &cobra.Command{
	Use:   "create-group",
	Short: "Creates a new group in the specified user pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_createGroupCmd).Standalone()

	cognitoIdp_createGroupCmd.Flags().String("description", "", "A description of the group that you're creating.")
	cognitoIdp_createGroupCmd.Flags().String("group-name", "", "A name for the group.")
	cognitoIdp_createGroupCmd.Flags().String("precedence", "", "A non-negative integer value that specifies the precedence of this group relative to the other groups that a user can belong to in the user pool.")
	cognitoIdp_createGroupCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) for the IAM role that you want to associate with the group.")
	cognitoIdp_createGroupCmd.Flags().String("user-pool-id", "", "The ID of the user pool where you want to create a user group.")
	cognitoIdp_createGroupCmd.MarkFlagRequired("group-name")
	cognitoIdp_createGroupCmd.MarkFlagRequired("user-pool-id")
	cognitoIdpCmd.AddCommand(cognitoIdp_createGroupCmd)
}
