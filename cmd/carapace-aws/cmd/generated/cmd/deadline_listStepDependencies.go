package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_listStepDependenciesCmd = &cobra.Command{
	Use:   "list-step-dependencies",
	Short: "Lists the dependencies for a step.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_listStepDependenciesCmd).Standalone()

	deadline_listStepDependenciesCmd.Flags().String("farm-id", "", "The farm ID for the step dependencies list.")
	deadline_listStepDependenciesCmd.Flags().String("job-id", "", "The job ID for the step dependencies list.")
	deadline_listStepDependenciesCmd.Flags().String("max-results", "", "The maximum number of results to return.")
	deadline_listStepDependenciesCmd.Flags().String("next-token", "", "The token for the next set of results, or `null` to start from the beginning.")
	deadline_listStepDependenciesCmd.Flags().String("queue-id", "", "The queue ID for the step dependencies list.")
	deadline_listStepDependenciesCmd.Flags().String("step-id", "", "The step ID to include on the list.")
	deadline_listStepDependenciesCmd.MarkFlagRequired("farm-id")
	deadline_listStepDependenciesCmd.MarkFlagRequired("job-id")
	deadline_listStepDependenciesCmd.MarkFlagRequired("queue-id")
	deadline_listStepDependenciesCmd.MarkFlagRequired("step-id")
	deadlineCmd.AddCommand(deadline_listStepDependenciesCmd)
}
