package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_listReferenceImportJobsCmd = &cobra.Command{
	Use:   "list-reference-import-jobs",
	Short: "Retrieves the metadata of one or more reference import jobs for a reference store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_listReferenceImportJobsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(omics_listReferenceImportJobsCmd).Standalone()

		omics_listReferenceImportJobsCmd.Flags().String("filter", "", "A filter to apply to the list.")
		omics_listReferenceImportJobsCmd.Flags().String("max-results", "", "The maximum number of jobs to return in one page of results.")
		omics_listReferenceImportJobsCmd.Flags().String("next-token", "", "Specify the pagination token from a previous request to retrieve the next page of results.")
		omics_listReferenceImportJobsCmd.Flags().String("reference-store-id", "", "The job's reference store ID.")
		omics_listReferenceImportJobsCmd.MarkFlagRequired("reference-store-id")
	})
	omicsCmd.AddCommand(omics_listReferenceImportJobsCmd)
}
