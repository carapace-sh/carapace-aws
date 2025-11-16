package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_listGroupMembershipsCmd = &cobra.Command{
	Use:   "list-group-memberships",
	Short: "Lists member users in a group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_listGroupMembershipsCmd).Standalone()

	quicksight_listGroupMembershipsCmd.Flags().String("aws-account-id", "", "The ID for the Amazon Web Services account that the group is in.")
	quicksight_listGroupMembershipsCmd.Flags().String("group-name", "", "The name of the group that you want to see a membership list of.")
	quicksight_listGroupMembershipsCmd.Flags().String("max-results", "", "The maximum number of results to return from this request.")
	quicksight_listGroupMembershipsCmd.Flags().String("namespace", "", "The namespace of the group that you want a list of users from.")
	quicksight_listGroupMembershipsCmd.Flags().String("next-token", "", "A pagination token that can be used in a subsequent request.")
	quicksight_listGroupMembershipsCmd.MarkFlagRequired("aws-account-id")
	quicksight_listGroupMembershipsCmd.MarkFlagRequired("group-name")
	quicksight_listGroupMembershipsCmd.MarkFlagRequired("namespace")
	quicksightCmd.AddCommand(quicksight_listGroupMembershipsCmd)
}
