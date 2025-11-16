package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_createGroupMembershipCmd = &cobra.Command{
	Use:   "create-group-membership",
	Short: "Adds an Amazon Quick Sight user to an Amazon Quick Sight group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_createGroupMembershipCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_createGroupMembershipCmd).Standalone()

		quicksight_createGroupMembershipCmd.Flags().String("aws-account-id", "", "The ID for the Amazon Web Services account that the group is in.")
		quicksight_createGroupMembershipCmd.Flags().String("group-name", "", "The name of the group that you want to add the user to.")
		quicksight_createGroupMembershipCmd.Flags().String("member-name", "", "The name of the user that you want to add to the group membership.")
		quicksight_createGroupMembershipCmd.Flags().String("namespace", "", "The namespace that you want the user to be a part of.")
		quicksight_createGroupMembershipCmd.MarkFlagRequired("aws-account-id")
		quicksight_createGroupMembershipCmd.MarkFlagRequired("group-name")
		quicksight_createGroupMembershipCmd.MarkFlagRequired("member-name")
		quicksight_createGroupMembershipCmd.MarkFlagRequired("namespace")
	})
	quicksightCmd.AddCommand(quicksight_createGroupMembershipCmd)
}
