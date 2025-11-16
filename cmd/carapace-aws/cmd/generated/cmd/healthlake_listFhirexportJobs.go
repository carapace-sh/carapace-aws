package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var healthlake_listFhirexportJobsCmd = &cobra.Command{
	Use:   "list-fhirexport-jobs",
	Short: "Lists all FHIR export jobs associated with an account and their statuses.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(healthlake_listFhirexportJobsCmd).Standalone()

	healthlake_listFhirexportJobsCmd.Flags().String("datastore-id", "", "Limits the response to the export job with the specified data store ID.")
	healthlake_listFhirexportJobsCmd.Flags().String("job-name", "", "Limits the response to the export job with the specified job name.")
	healthlake_listFhirexportJobsCmd.Flags().String("job-status", "", "Limits the response to export jobs with the specified job status.")
	healthlake_listFhirexportJobsCmd.Flags().String("max-results", "", "Limits the number of results returned for a ListFHIRExportJobs to a maximum quantity specified by the user.")
	healthlake_listFhirexportJobsCmd.Flags().String("next-token", "", "A pagination token used to identify the next page of results to return.")
	healthlake_listFhirexportJobsCmd.Flags().String("submitted-after", "", "Limits the response to FHIR export jobs submitted after a user-specified date.")
	healthlake_listFhirexportJobsCmd.Flags().String("submitted-before", "", "Limits the response to FHIR export jobs submitted before a user- specified date.")
	healthlake_listFhirexportJobsCmd.MarkFlagRequired("datastore-id")
	healthlakeCmd.AddCommand(healthlake_listFhirexportJobsCmd)
}
