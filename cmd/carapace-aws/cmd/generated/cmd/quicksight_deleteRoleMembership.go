package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_deleteRoleMembershipCmd = &cobra.Command{
	Use:   "delete-role-membership",
	Short: "Removes a group from a role.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_deleteRoleMembershipCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_deleteRoleMembershipCmd).Standalone()

		quicksight_deleteRoleMembershipCmd.Flags().String("aws-account-id", "", "The ID for the Amazon Web Services account that you want to create a group in.")
		quicksight_deleteRoleMembershipCmd.Flags().String("member-name", "", "The name of the group.")
		quicksight_deleteRoleMembershipCmd.Flags().String("namespace", "", "The namespace that contains the role.")
		quicksight_deleteRoleMembershipCmd.Flags().String("role", "", "The role that you want to remove permissions from.")
		quicksight_deleteRoleMembershipCmd.MarkFlagRequired("aws-account-id")
		quicksight_deleteRoleMembershipCmd.MarkFlagRequired("member-name")
		quicksight_deleteRoleMembershipCmd.MarkFlagRequired("namespace")
		quicksight_deleteRoleMembershipCmd.MarkFlagRequired("role")
	})
	quicksightCmd.AddCommand(quicksight_deleteRoleMembershipCmd)
}
