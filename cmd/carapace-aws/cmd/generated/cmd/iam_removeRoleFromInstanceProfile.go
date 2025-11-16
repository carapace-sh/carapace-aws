package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_removeRoleFromInstanceProfileCmd = &cobra.Command{
	Use:   "remove-role-from-instance-profile",
	Short: "Removes the specified IAM role from the specified Amazon EC2 instance profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_removeRoleFromInstanceProfileCmd).Standalone()

	iam_removeRoleFromInstanceProfileCmd.Flags().String("instance-profile-name", "", "The name of the instance profile to update.")
	iam_removeRoleFromInstanceProfileCmd.Flags().String("role-name", "", "The name of the role to remove.")
	iam_removeRoleFromInstanceProfileCmd.MarkFlagRequired("instance-profile-name")
	iam_removeRoleFromInstanceProfileCmd.MarkFlagRequired("role-name")
	iamCmd.AddCommand(iam_removeRoleFromInstanceProfileCmd)
}
