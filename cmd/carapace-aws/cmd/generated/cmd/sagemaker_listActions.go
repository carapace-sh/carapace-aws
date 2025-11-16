package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listActionsCmd = &cobra.Command{
	Use:   "list-actions",
	Short: "Lists the actions in your account and their properties.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listActionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_listActionsCmd).Standalone()

		sagemaker_listActionsCmd.Flags().String("action-type", "", "A filter that returns only actions of the specified type.")
		sagemaker_listActionsCmd.Flags().String("created-after", "", "A filter that returns only actions created on or after the specified time.")
		sagemaker_listActionsCmd.Flags().String("created-before", "", "A filter that returns only actions created on or before the specified time.")
		sagemaker_listActionsCmd.Flags().String("max-results", "", "The maximum number of actions to return in the response.")
		sagemaker_listActionsCmd.Flags().String("next-token", "", "If the previous call to `ListActions` didn't return the full set of actions, the call returns a token for getting the next set of actions.")
		sagemaker_listActionsCmd.Flags().String("sort-by", "", "The property used to sort results.")
		sagemaker_listActionsCmd.Flags().String("sort-order", "", "The sort order.")
		sagemaker_listActionsCmd.Flags().String("source-uri", "", "A filter that returns only actions with the specified source URI.")
	})
	sagemakerCmd.AddCommand(sagemaker_listActionsCmd)
}
