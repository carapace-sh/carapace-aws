package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_listDataDeletionJobsCmd = &cobra.Command{
	Use:   "list-data-deletion-jobs",
	Short: "Returns a list of data deletion jobs for a dataset group ordered by creation time, with the most recent first.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_listDataDeletionJobsCmd).Standalone()

	personalize_listDataDeletionJobsCmd.Flags().String("dataset-group-arn", "", "The Amazon Resource Name (ARN) of the dataset group to list data deletion jobs for.")
	personalize_listDataDeletionJobsCmd.Flags().String("max-results", "", "The maximum number of data deletion jobs to return.")
	personalize_listDataDeletionJobsCmd.Flags().String("next-token", "", "A token returned from the previous call to `ListDataDeletionJobs` for getting the next set of jobs (if they exist).")
	personalizeCmd.AddCommand(personalize_listDataDeletionJobsCmd)
}
