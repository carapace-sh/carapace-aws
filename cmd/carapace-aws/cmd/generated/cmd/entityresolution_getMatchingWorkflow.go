package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var entityresolution_getMatchingWorkflowCmd = &cobra.Command{
	Use:   "get-matching-workflow",
	Short: "Returns the `MatchingWorkflow` with a given name, if it exists.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(entityresolution_getMatchingWorkflowCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(entityresolution_getMatchingWorkflowCmd).Standalone()

		entityresolution_getMatchingWorkflowCmd.Flags().String("workflow-name", "", "The name of the workflow.")
		entityresolution_getMatchingWorkflowCmd.MarkFlagRequired("workflow-name")
	})
	entityresolutionCmd.AddCommand(entityresolution_getMatchingWorkflowCmd)
}
