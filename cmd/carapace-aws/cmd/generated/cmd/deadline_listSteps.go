package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_listStepsCmd = &cobra.Command{
	Use:   "list-steps",
	Short: "Lists steps for a job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_listStepsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(deadline_listStepsCmd).Standalone()

		deadline_listStepsCmd.Flags().String("farm-id", "", "The farm ID to include on the list of steps.")
		deadline_listStepsCmd.Flags().String("job-id", "", "The job ID to include on the list of steps.")
		deadline_listStepsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
		deadline_listStepsCmd.Flags().String("next-token", "", "The token for the next set of results, or `null` to start from the beginning.")
		deadline_listStepsCmd.Flags().String("queue-id", "", "The queue ID to include on the list of steps.")
		deadline_listStepsCmd.MarkFlagRequired("farm-id")
		deadline_listStepsCmd.MarkFlagRequired("job-id")
		deadline_listStepsCmd.MarkFlagRequired("queue-id")
	})
	deadlineCmd.AddCommand(deadline_listStepsCmd)
}
