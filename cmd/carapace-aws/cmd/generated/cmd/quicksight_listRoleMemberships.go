package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_listRoleMembershipsCmd = &cobra.Command{
	Use:   "list-role-memberships",
	Short: "Lists all groups that are associated with a role.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_listRoleMembershipsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_listRoleMembershipsCmd).Standalone()

		quicksight_listRoleMembershipsCmd.Flags().String("aws-account-id", "", "The ID for the Amazon Web Services account that you want to create a group in.")
		quicksight_listRoleMembershipsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
		quicksight_listRoleMembershipsCmd.Flags().String("namespace", "", "The namespace that includes the role.")
		quicksight_listRoleMembershipsCmd.Flags().String("next-token", "", "A pagination token that can be used in a subsequent request.")
		quicksight_listRoleMembershipsCmd.Flags().String("role", "", "The name of the role.")
		quicksight_listRoleMembershipsCmd.MarkFlagRequired("aws-account-id")
		quicksight_listRoleMembershipsCmd.MarkFlagRequired("namespace")
		quicksight_listRoleMembershipsCmd.MarkFlagRequired("role")
	})
	quicksightCmd.AddCommand(quicksight_listRoleMembershipsCmd)
}
