package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_createRoleMembershipCmd = &cobra.Command{
	Use:   "create-role-membership",
	Short: "Use `CreateRoleMembership` to add an existing Quick Sight group to an existing role.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_createRoleMembershipCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_createRoleMembershipCmd).Standalone()

		quicksight_createRoleMembershipCmd.Flags().String("aws-account-id", "", "The ID for the Amazon Web Services account that you want to create a group in.")
		quicksight_createRoleMembershipCmd.Flags().String("member-name", "", "The name of the group that you want to add to the role.")
		quicksight_createRoleMembershipCmd.Flags().String("namespace", "", "The namespace that the role belongs to.")
		quicksight_createRoleMembershipCmd.Flags().String("role", "", "The role that you want to add a group to.")
		quicksight_createRoleMembershipCmd.MarkFlagRequired("aws-account-id")
		quicksight_createRoleMembershipCmd.MarkFlagRequired("member-name")
		quicksight_createRoleMembershipCmd.MarkFlagRequired("namespace")
		quicksight_createRoleMembershipCmd.MarkFlagRequired("role")
	})
	quicksightCmd.AddCommand(quicksight_createRoleMembershipCmd)
}
