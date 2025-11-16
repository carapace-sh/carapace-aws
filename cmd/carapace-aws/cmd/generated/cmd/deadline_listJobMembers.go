package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_listJobMembersCmd = &cobra.Command{
	Use:   "list-job-members",
	Short: "Lists members on a job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_listJobMembersCmd).Standalone()

	deadline_listJobMembersCmd.Flags().String("farm-id", "", "The farm ID of the job to list.")
	deadline_listJobMembersCmd.Flags().String("job-id", "", "The job ID to include on the list.")
	deadline_listJobMembersCmd.Flags().String("max-results", "", "The maximum number of results to return.")
	deadline_listJobMembersCmd.Flags().String("next-token", "", "The token for the next set of results, or `null` to start from the beginning.")
	deadline_listJobMembersCmd.Flags().String("queue-id", "", "The queue ID to include on the list.")
	deadline_listJobMembersCmd.MarkFlagRequired("farm-id")
	deadline_listJobMembersCmd.MarkFlagRequired("job-id")
	deadline_listJobMembersCmd.MarkFlagRequired("queue-id")
	deadlineCmd.AddCommand(deadline_listJobMembersCmd)
}
