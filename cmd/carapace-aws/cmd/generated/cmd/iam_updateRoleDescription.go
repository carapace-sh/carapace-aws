package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_updateRoleDescriptionCmd = &cobra.Command{
	Use:   "update-role-description",
	Short: "Use [UpdateRole](https://docs.aws.amazon.com/IAM/latest/APIReference/API_UpdateRole.html) instead.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_updateRoleDescriptionCmd).Standalone()

	iam_updateRoleDescriptionCmd.Flags().String("description", "", "The new description that you want to apply to the specified role.")
	iam_updateRoleDescriptionCmd.Flags().String("role-name", "", "The name of the role that you want to modify.")
	iam_updateRoleDescriptionCmd.MarkFlagRequired("description")
	iam_updateRoleDescriptionCmd.MarkFlagRequired("role-name")
	iamCmd.AddCommand(iam_updateRoleDescriptionCmd)
}
