package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_listUserGroupsCmd = &cobra.Command{
	Use:   "list-user-groups",
	Short: "Lists the Amazon Quick Sight groups that an Amazon Quick Sight user is a member of.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_listUserGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_listUserGroupsCmd).Standalone()

		quicksight_listUserGroupsCmd.Flags().String("aws-account-id", "", "The Amazon Web Services account ID that the user is in.")
		quicksight_listUserGroupsCmd.Flags().String("max-results", "", "The maximum number of results to return from this request.")
		quicksight_listUserGroupsCmd.Flags().String("namespace", "", "The namespace.")
		quicksight_listUserGroupsCmd.Flags().String("next-token", "", "A pagination token that can be used in a subsequent request.")
		quicksight_listUserGroupsCmd.Flags().String("user-name", "", "The Amazon Quick Sight user name that you want to list group memberships for.")
		quicksight_listUserGroupsCmd.MarkFlagRequired("aws-account-id")
		quicksight_listUserGroupsCmd.MarkFlagRequired("namespace")
		quicksight_listUserGroupsCmd.MarkFlagRequired("user-name")
	})
	quicksightCmd.AddCommand(quicksight_listUserGroupsCmd)
}
