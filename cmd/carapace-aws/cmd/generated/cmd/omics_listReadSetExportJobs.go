package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_listReadSetExportJobsCmd = &cobra.Command{
	Use:   "list-read-set-export-jobs",
	Short: "Retrieves a list of read set export jobs in a JSON formatted response.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_listReadSetExportJobsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(omics_listReadSetExportJobsCmd).Standalone()

		omics_listReadSetExportJobsCmd.Flags().String("filter", "", "A filter to apply to the list.")
		omics_listReadSetExportJobsCmd.Flags().String("max-results", "", "The maximum number of jobs to return in one page of results.")
		omics_listReadSetExportJobsCmd.Flags().String("next-token", "", "Specify the pagination token from a previous request to retrieve the next page of results.")
		omics_listReadSetExportJobsCmd.Flags().String("sequence-store-id", "", "The jobs' sequence store ID.")
		omics_listReadSetExportJobsCmd.MarkFlagRequired("sequence-store-id")
	})
	omicsCmd.AddCommand(omics_listReadSetExportJobsCmd)
}
