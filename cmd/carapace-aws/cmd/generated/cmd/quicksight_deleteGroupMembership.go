package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_deleteGroupMembershipCmd = &cobra.Command{
	Use:   "delete-group-membership",
	Short: "Removes a user from a group so that the user is no longer a member of the group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_deleteGroupMembershipCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_deleteGroupMembershipCmd).Standalone()

		quicksight_deleteGroupMembershipCmd.Flags().String("aws-account-id", "", "The ID for the Amazon Web Services account that the group is in.")
		quicksight_deleteGroupMembershipCmd.Flags().String("group-name", "", "The name of the group that you want to delete the user from.")
		quicksight_deleteGroupMembershipCmd.Flags().String("member-name", "", "The name of the user that you want to delete from the group membership.")
		quicksight_deleteGroupMembershipCmd.Flags().String("namespace", "", "The namespace of the group that you want to remove a user from.")
		quicksight_deleteGroupMembershipCmd.MarkFlagRequired("aws-account-id")
		quicksight_deleteGroupMembershipCmd.MarkFlagRequired("group-name")
		quicksight_deleteGroupMembershipCmd.MarkFlagRequired("member-name")
		quicksight_deleteGroupMembershipCmd.MarkFlagRequired("namespace")
	})
	quicksightCmd.AddCommand(quicksight_deleteGroupMembershipCmd)
}
