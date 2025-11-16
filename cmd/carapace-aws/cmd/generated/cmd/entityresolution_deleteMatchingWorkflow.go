package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var entityresolution_deleteMatchingWorkflowCmd = &cobra.Command{
	Use:   "delete-matching-workflow",
	Short: "Deletes the `MatchingWorkflow` with a given name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(entityresolution_deleteMatchingWorkflowCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(entityresolution_deleteMatchingWorkflowCmd).Standalone()

		entityresolution_deleteMatchingWorkflowCmd.Flags().String("workflow-name", "", "The name of the workflow to be retrieved.")
		entityresolution_deleteMatchingWorkflowCmd.MarkFlagRequired("workflow-name")
	})
	entityresolutionCmd.AddCommand(entityresolution_deleteMatchingWorkflowCmd)
}
