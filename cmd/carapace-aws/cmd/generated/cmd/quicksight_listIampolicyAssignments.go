package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_listIampolicyAssignmentsCmd = &cobra.Command{
	Use:   "list-iampolicy-assignments",
	Short: "Lists the IAM policy assignments in the current Amazon Quick Sight account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_listIampolicyAssignmentsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_listIampolicyAssignmentsCmd).Standalone()

		quicksight_listIampolicyAssignmentsCmd.Flags().String("assignment-status", "", "The status of the assignments.")
		quicksight_listIampolicyAssignmentsCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains these IAM policy assignments.")
		quicksight_listIampolicyAssignmentsCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
		quicksight_listIampolicyAssignmentsCmd.Flags().String("namespace", "", "The namespace for the assignments.")
		quicksight_listIampolicyAssignmentsCmd.Flags().String("next-token", "", "The token for the next set of results, or null if there are no more results.")
		quicksight_listIampolicyAssignmentsCmd.MarkFlagRequired("aws-account-id")
		quicksight_listIampolicyAssignmentsCmd.MarkFlagRequired("namespace")
	})
	quicksightCmd.AddCommand(quicksight_listIampolicyAssignmentsCmd)
}
