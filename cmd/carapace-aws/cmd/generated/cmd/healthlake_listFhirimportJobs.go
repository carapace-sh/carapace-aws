package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var healthlake_listFhirimportJobsCmd = &cobra.Command{
	Use:   "list-fhirimport-jobs",
	Short: "List all FHIR import jobs associated with an account and their statuses.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(healthlake_listFhirimportJobsCmd).Standalone()

	healthlake_listFhirimportJobsCmd.Flags().String("datastore-id", "", "Limits the response to the import job with the specified data store ID.")
	healthlake_listFhirimportJobsCmd.Flags().String("job-name", "", "Limits the response to the import job with the specified job name.")
	healthlake_listFhirimportJobsCmd.Flags().String("job-status", "", "Limits the response to the import job with the specified job status.")
	healthlake_listFhirimportJobsCmd.Flags().String("max-results", "", "Limits the number of results returned for `ListFHIRImportJobs` to a maximum quantity specified by the user.")
	healthlake_listFhirimportJobsCmd.Flags().String("next-token", "", "The pagination token used to identify the next page of results to return.")
	healthlake_listFhirimportJobsCmd.Flags().String("submitted-after", "", "Limits the response to FHIR import jobs submitted after a user-specified date.")
	healthlake_listFhirimportJobsCmd.Flags().String("submitted-before", "", "Limits the response to FHIR import jobs submitted before a user- specified date.")
	healthlake_listFhirimportJobsCmd.MarkFlagRequired("datastore-id")
	healthlakeCmd.AddCommand(healthlake_listFhirimportJobsCmd)
}
