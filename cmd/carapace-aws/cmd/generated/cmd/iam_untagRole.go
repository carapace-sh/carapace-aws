package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_untagRoleCmd = &cobra.Command{
	Use:   "untag-role",
	Short: "Removes the specified tags from the role.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_untagRoleCmd).Standalone()

	iam_untagRoleCmd.Flags().String("role-name", "", "The name of the IAM role from which you want to remove tags.")
	iam_untagRoleCmd.Flags().String("tag-keys", "", "A list of key names as a simple array of strings.")
	iam_untagRoleCmd.MarkFlagRequired("role-name")
	iam_untagRoleCmd.MarkFlagRequired("tag-keys")
	iamCmd.AddCommand(iam_untagRoleCmd)
}
