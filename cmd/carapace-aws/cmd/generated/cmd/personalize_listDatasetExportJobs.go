package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_listDatasetExportJobsCmd = &cobra.Command{
	Use:   "list-dataset-export-jobs",
	Short: "Returns a list of dataset export jobs that use the given dataset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_listDatasetExportJobsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(personalize_listDatasetExportJobsCmd).Standalone()

		personalize_listDatasetExportJobsCmd.Flags().String("dataset-arn", "", "The Amazon Resource Name (ARN) of the dataset to list the dataset export jobs for.")
		personalize_listDatasetExportJobsCmd.Flags().String("max-results", "", "The maximum number of dataset export jobs to return.")
		personalize_listDatasetExportJobsCmd.Flags().String("next-token", "", "A token returned from the previous call to `ListDatasetExportJobs` for getting the next set of dataset export jobs (if they exist).")
	})
	personalizeCmd.AddCommand(personalize_listDatasetExportJobsCmd)
}
