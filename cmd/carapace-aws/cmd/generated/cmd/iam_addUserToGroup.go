package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_addUserToGroupCmd = &cobra.Command{
	Use:   "add-user-to-group",
	Short: "Adds the specified user to the specified group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_addUserToGroupCmd).Standalone()

	iam_addUserToGroupCmd.Flags().String("group-name", "", "The name of the group to update.")
	iam_addUserToGroupCmd.Flags().String("user-name", "", "The name of the user to add.")
	iam_addUserToGroupCmd.MarkFlagRequired("group-name")
	iam_addUserToGroupCmd.MarkFlagRequired("user-name")
	iamCmd.AddCommand(iam_addUserToGroupCmd)
}
