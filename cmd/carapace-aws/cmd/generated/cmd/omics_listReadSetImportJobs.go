package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_listReadSetImportJobsCmd = &cobra.Command{
	Use:   "list-read-set-import-jobs",
	Short: "Retrieves a list of read set import jobs and returns the data in JSON format.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_listReadSetImportJobsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(omics_listReadSetImportJobsCmd).Standalone()

		omics_listReadSetImportJobsCmd.Flags().String("filter", "", "A filter to apply to the list.")
		omics_listReadSetImportJobsCmd.Flags().String("max-results", "", "The maximum number of jobs to return in one page of results.")
		omics_listReadSetImportJobsCmd.Flags().String("next-token", "", "Specify the pagination token from a previous request to retrieve the next page of results.")
		omics_listReadSetImportJobsCmd.Flags().String("sequence-store-id", "", "The jobs' sequence store ID.")
		omics_listReadSetImportJobsCmd.MarkFlagRequired("sequence-store-id")
	})
	omicsCmd.AddCommand(omics_listReadSetImportJobsCmd)
}
