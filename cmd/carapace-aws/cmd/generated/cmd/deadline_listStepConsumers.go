package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_listStepConsumersCmd = &cobra.Command{
	Use:   "list-step-consumers",
	Short: "Lists step consumers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_listStepConsumersCmd).Standalone()

	deadline_listStepConsumersCmd.Flags().String("farm-id", "", "The farm ID for the list of step consumers.")
	deadline_listStepConsumersCmd.Flags().String("job-id", "", "The job ID for the step consumer.")
	deadline_listStepConsumersCmd.Flags().String("max-results", "", "The maximum number of results to return.")
	deadline_listStepConsumersCmd.Flags().String("next-token", "", "The token for the next set of results, or `null` to start from the beginning.")
	deadline_listStepConsumersCmd.Flags().String("queue-id", "", "The queue ID for the step consumer.")
	deadline_listStepConsumersCmd.Flags().String("step-id", "", "The step ID to include on the list.")
	deadline_listStepConsumersCmd.MarkFlagRequired("farm-id")
	deadline_listStepConsumersCmd.MarkFlagRequired("job-id")
	deadline_listStepConsumersCmd.MarkFlagRequired("queue-id")
	deadline_listStepConsumersCmd.MarkFlagRequired("step-id")
	deadlineCmd.AddCommand(deadline_listStepConsumersCmd)
}
