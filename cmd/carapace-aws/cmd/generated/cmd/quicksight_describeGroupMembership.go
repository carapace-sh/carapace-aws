package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_describeGroupMembershipCmd = &cobra.Command{
	Use:   "describe-group-membership",
	Short: "Use the `DescribeGroupMembership` operation to determine if a user is a member of the specified group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_describeGroupMembershipCmd).Standalone()

	quicksight_describeGroupMembershipCmd.Flags().String("aws-account-id", "", "The ID for the Amazon Web Services account that the group is in.")
	quicksight_describeGroupMembershipCmd.Flags().String("group-name", "", "The name of the group that you want to search.")
	quicksight_describeGroupMembershipCmd.Flags().String("member-name", "", "The user name of the user that you want to search for.")
	quicksight_describeGroupMembershipCmd.Flags().String("namespace", "", "The namespace that includes the group you are searching within.")
	quicksight_describeGroupMembershipCmd.MarkFlagRequired("aws-account-id")
	quicksight_describeGroupMembershipCmd.MarkFlagRequired("group-name")
	quicksight_describeGroupMembershipCmd.MarkFlagRequired("member-name")
	quicksight_describeGroupMembershipCmd.MarkFlagRequired("namespace")
	quicksightCmd.AddCommand(quicksight_describeGroupMembershipCmd)
}
