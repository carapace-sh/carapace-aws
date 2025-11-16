package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_listVariantImportJobsCmd = &cobra.Command{
	Use:   "list-variant-import-jobs",
	Short: "Amazon Web Services HealthOmics variant stores and annotation stores will no longer be open to new customers starting November 7, 2025. If you would like to use variant stores or annotation stores, sign up prior to that date.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_listVariantImportJobsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(omics_listVariantImportJobsCmd).Standalone()

		omics_listVariantImportJobsCmd.Flags().String("filter", "", "A filter to apply to the list.")
		omics_listVariantImportJobsCmd.Flags().String("ids", "", "A list of job IDs.")
		omics_listVariantImportJobsCmd.Flags().String("max-results", "", "The maximum number of import jobs to return in one page of results.")
		omics_listVariantImportJobsCmd.Flags().String("next-token", "", "Specify the pagination token from a previous request to retrieve the next page of results.")
	})
	omicsCmd.AddCommand(omics_listVariantImportJobsCmd)
}
