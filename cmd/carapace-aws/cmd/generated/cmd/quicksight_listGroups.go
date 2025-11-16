package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_listGroupsCmd = &cobra.Command{
	Use:   "list-groups",
	Short: "Lists all user groups in Amazon Quick Sight.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_listGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_listGroupsCmd).Standalone()

		quicksight_listGroupsCmd.Flags().String("aws-account-id", "", "The ID for the Amazon Web Services account that the group is in.")
		quicksight_listGroupsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
		quicksight_listGroupsCmd.Flags().String("namespace", "", "The namespace that you want a list of groups from.")
		quicksight_listGroupsCmd.Flags().String("next-token", "", "A pagination token that can be used in a subsequent request.")
		quicksight_listGroupsCmd.MarkFlagRequired("aws-account-id")
		quicksight_listGroupsCmd.MarkFlagRequired("namespace")
	})
	quicksightCmd.AddCommand(quicksight_listGroupsCmd)
}
