package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_addRoleToInstanceProfileCmd = &cobra.Command{
	Use:   "add-role-to-instance-profile",
	Short: "Adds the specified IAM role to the specified instance profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_addRoleToInstanceProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_addRoleToInstanceProfileCmd).Standalone()

		iam_addRoleToInstanceProfileCmd.Flags().String("instance-profile-name", "", "The name of the instance profile to update.")
		iam_addRoleToInstanceProfileCmd.Flags().String("role-name", "", "The name of the role to add.")
		iam_addRoleToInstanceProfileCmd.MarkFlagRequired("instance-profile-name")
		iam_addRoleToInstanceProfileCmd.MarkFlagRequired("role-name")
	})
	iamCmd.AddCommand(iam_addRoleToInstanceProfileCmd)
}
