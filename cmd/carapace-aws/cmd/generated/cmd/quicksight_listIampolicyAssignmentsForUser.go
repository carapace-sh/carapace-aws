package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_listIampolicyAssignmentsForUserCmd = &cobra.Command{
	Use:   "list-iampolicy-assignments-for-user",
	Short: "Lists all of the IAM policy assignments, including the Amazon Resource Names (ARNs), for the IAM policies assigned to the specified user and group, or groups that the user belongs to.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_listIampolicyAssignmentsForUserCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_listIampolicyAssignmentsForUserCmd).Standalone()

		quicksight_listIampolicyAssignmentsForUserCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the assignments.")
		quicksight_listIampolicyAssignmentsForUserCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
		quicksight_listIampolicyAssignmentsForUserCmd.Flags().String("namespace", "", "The namespace of the assignment.")
		quicksight_listIampolicyAssignmentsForUserCmd.Flags().String("next-token", "", "The token for the next set of results, or null if there are no more results.")
		quicksight_listIampolicyAssignmentsForUserCmd.Flags().String("user-name", "", "The name of the user.")
		quicksight_listIampolicyAssignmentsForUserCmd.MarkFlagRequired("aws-account-id")
		quicksight_listIampolicyAssignmentsForUserCmd.MarkFlagRequired("namespace")
		quicksight_listIampolicyAssignmentsForUserCmd.MarkFlagRequired("user-name")
	})
	quicksightCmd.AddCommand(quicksight_listIampolicyAssignmentsForUserCmd)
}
