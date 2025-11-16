package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_updateGroupCmd = &cobra.Command{
	Use:   "update-group",
	Short: "Updates the name and/or the path of the specified IAM group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_updateGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_updateGroupCmd).Standalone()

		iam_updateGroupCmd.Flags().String("group-name", "", "Name of the IAM group to update.")
		iam_updateGroupCmd.Flags().String("new-group-name", "", "New name for the IAM group.")
		iam_updateGroupCmd.Flags().String("new-path", "", "New path for the IAM group.")
		iam_updateGroupCmd.MarkFlagRequired("group-name")
	})
	iamCmd.AddCommand(iam_updateGroupCmd)
}
