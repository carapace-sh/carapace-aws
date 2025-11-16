package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_listDatasetImportJobsCmd = &cobra.Command{
	Use:   "list-dataset-import-jobs",
	Short: "Returns a list of dataset import jobs that use the given dataset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_listDatasetImportJobsCmd).Standalone()

	personalize_listDatasetImportJobsCmd.Flags().String("dataset-arn", "", "The Amazon Resource Name (ARN) of the dataset to list the dataset import jobs for.")
	personalize_listDatasetImportJobsCmd.Flags().String("max-results", "", "The maximum number of dataset import jobs to return.")
	personalize_listDatasetImportJobsCmd.Flags().String("next-token", "", "A token returned from the previous call to `ListDatasetImportJobs` for getting the next set of dataset import jobs (if they exist).")
	personalizeCmd.AddCommand(personalize_listDatasetImportJobsCmd)
}
