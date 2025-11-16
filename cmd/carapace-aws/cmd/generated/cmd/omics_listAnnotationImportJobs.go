package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_listAnnotationImportJobsCmd = &cobra.Command{
	Use:   "list-annotation-import-jobs",
	Short: "Amazon Web Services HealthOmics variant stores and annotation stores will no longer be open to new customers starting November 7, 2025. If you would like to use variant stores or annotation stores, sign up prior to that date.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_listAnnotationImportJobsCmd).Standalone()

	omics_listAnnotationImportJobsCmd.Flags().String("filter", "", "A filter to apply to the list.")
	omics_listAnnotationImportJobsCmd.Flags().String("ids", "", "IDs of annotation import jobs to retrieve.")
	omics_listAnnotationImportJobsCmd.Flags().String("max-results", "", "The maximum number of jobs to return in one page of results.")
	omics_listAnnotationImportJobsCmd.Flags().String("next-token", "", "Specifies the pagination token from a previous request to retrieve the next page of results.")
	omicsCmd.AddCommand(omics_listAnnotationImportJobsCmd)
}
