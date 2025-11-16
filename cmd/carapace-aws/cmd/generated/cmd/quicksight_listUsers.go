package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_listUsersCmd = &cobra.Command{
	Use:   "list-users",
	Short: "Returns a list of all of the Amazon Quick Sight users belonging to this account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_listUsersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_listUsersCmd).Standalone()

		quicksight_listUsersCmd.Flags().String("aws-account-id", "", "The ID for the Amazon Web Services account that the user is in.")
		quicksight_listUsersCmd.Flags().String("max-results", "", "The maximum number of results to return from this request.")
		quicksight_listUsersCmd.Flags().String("namespace", "", "The namespace.")
		quicksight_listUsersCmd.Flags().String("next-token", "", "A pagination token that can be used in a subsequent request.")
		quicksight_listUsersCmd.MarkFlagRequired("aws-account-id")
		quicksight_listUsersCmd.MarkFlagRequired("namespace")
	})
	quicksightCmd.AddCommand(quicksight_listUsersCmd)
}
