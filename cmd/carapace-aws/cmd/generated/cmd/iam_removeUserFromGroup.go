package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_removeUserFromGroupCmd = &cobra.Command{
	Use:   "remove-user-from-group",
	Short: "Removes the specified user from the specified group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_removeUserFromGroupCmd).Standalone()

	iam_removeUserFromGroupCmd.Flags().String("group-name", "", "The name of the group to update.")
	iam_removeUserFromGroupCmd.Flags().String("user-name", "", "The name of the user to remove.")
	iam_removeUserFromGroupCmd.MarkFlagRequired("group-name")
	iam_removeUserFromGroupCmd.MarkFlagRequired("user-name")
	iamCmd.AddCommand(iam_removeUserFromGroupCmd)
}
