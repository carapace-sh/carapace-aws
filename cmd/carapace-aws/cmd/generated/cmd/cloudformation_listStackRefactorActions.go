package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_listStackRefactorActionsCmd = &cobra.Command{
	Use:   "list-stack-refactor-actions",
	Short: "Lists the stack refactor actions that will be taken after calling the [ExecuteStackRefactor]() action.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_listStackRefactorActionsCmd).Standalone()

	cloudformation_listStackRefactorActionsCmd.Flags().String("max-results", "", "The maximum number of results to be returned with a single call.")
	cloudformation_listStackRefactorActionsCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
	cloudformation_listStackRefactorActionsCmd.Flags().String("stack-refactor-id", "", "The ID associated with the stack refactor created from the [CreateStackRefactor]() action.")
	cloudformation_listStackRefactorActionsCmd.MarkFlagRequired("stack-refactor-id")
	cloudformationCmd.AddCommand(cloudformation_listStackRefactorActionsCmd)
}
