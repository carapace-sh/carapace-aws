package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_deleteGroupCmd = &cobra.Command{
	Use:   "delete-group",
	Short: "Deletes the specified IAM group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_deleteGroupCmd).Standalone()

	iam_deleteGroupCmd.Flags().String("group-name", "", "The name of the IAM group to delete.")
	iam_deleteGroupCmd.MarkFlagRequired("group-name")
	iamCmd.AddCommand(iam_deleteGroupCmd)
}
