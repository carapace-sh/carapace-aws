package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_deleteRoleCmd = &cobra.Command{
	Use:   "delete-role",
	Short: "Deletes the specified role.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_deleteRoleCmd).Standalone()

	iam_deleteRoleCmd.Flags().String("role-name", "", "The name of the role to delete.")
	iam_deleteRoleCmd.MarkFlagRequired("role-name")
	iamCmd.AddCommand(iam_deleteRoleCmd)
}
