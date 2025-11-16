package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_listSessionsForWorkerCmd = &cobra.Command{
	Use:   "list-sessions-for-worker",
	Short: "Lists sessions for a worker.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_listSessionsForWorkerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(deadline_listSessionsForWorkerCmd).Standalone()

		deadline_listSessionsForWorkerCmd.Flags().String("farm-id", "", "The farm ID for the session.")
		deadline_listSessionsForWorkerCmd.Flags().String("fleet-id", "", "The fleet ID for the session.")
		deadline_listSessionsForWorkerCmd.Flags().String("max-results", "", "The maximum number of results to return.")
		deadline_listSessionsForWorkerCmd.Flags().String("next-token", "", "The token for the next set of results, or `null` to start from the beginning.")
		deadline_listSessionsForWorkerCmd.Flags().String("worker-id", "", "The worker ID for the session.")
		deadline_listSessionsForWorkerCmd.MarkFlagRequired("farm-id")
		deadline_listSessionsForWorkerCmd.MarkFlagRequired("fleet-id")
		deadline_listSessionsForWorkerCmd.MarkFlagRequired("worker-id")
	})
	deadlineCmd.AddCommand(deadline_listSessionsForWorkerCmd)
}
