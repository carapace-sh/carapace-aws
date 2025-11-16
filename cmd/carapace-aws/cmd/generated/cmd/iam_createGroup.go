package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_createGroupCmd = &cobra.Command{
	Use:   "create-group",
	Short: "Creates a new group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_createGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_createGroupCmd).Standalone()

		iam_createGroupCmd.Flags().String("group-name", "", "The name of the group to create.")
		iam_createGroupCmd.Flags().String("path", "", "The path to the group.")
		iam_createGroupCmd.MarkFlagRequired("group-name")
	})
	iamCmd.AddCommand(iam_createGroupCmd)
}
