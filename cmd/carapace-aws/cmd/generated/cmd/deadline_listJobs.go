package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_listJobsCmd = &cobra.Command{
	Use:   "list-jobs",
	Short: "Lists jobs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_listJobsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(deadline_listJobsCmd).Standalone()

		deadline_listJobsCmd.Flags().String("farm-id", "", "The farm ID for the jobs.")
		deadline_listJobsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
		deadline_listJobsCmd.Flags().String("next-token", "", "The token for the next set of results, or `null` to start from the beginning.")
		deadline_listJobsCmd.Flags().String("principal-id", "", "The principal ID of the members on the jobs.")
		deadline_listJobsCmd.Flags().String("queue-id", "", "The queue ID for the job.")
		deadline_listJobsCmd.MarkFlagRequired("farm-id")
		deadline_listJobsCmd.MarkFlagRequired("queue-id")
	})
	deadlineCmd.AddCommand(deadline_listJobsCmd)
}
