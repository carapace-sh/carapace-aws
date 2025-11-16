package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transcribe_listMedicalScribeJobsCmd = &cobra.Command{
	Use:   "list-medical-scribe-jobs",
	Short: "Provides a list of Medical Scribe jobs that match the specified criteria.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transcribe_listMedicalScribeJobsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(transcribe_listMedicalScribeJobsCmd).Standalone()

		transcribe_listMedicalScribeJobsCmd.Flags().String("job-name-contains", "", "Returns only the Medical Scribe jobs that contain the specified string.")
		transcribe_listMedicalScribeJobsCmd.Flags().String("max-results", "", "The maximum number of Medical Scribe jobs to return in each page of results.")
		transcribe_listMedicalScribeJobsCmd.Flags().String("next-token", "", "If your `ListMedicalScribeJobs` request returns more results than can be displayed, `NextToken` is displayed in the response with an associated string.")
		transcribe_listMedicalScribeJobsCmd.Flags().String("status", "", "Returns only Medical Scribe jobs with the specified status.")
	})
	transcribeCmd.AddCommand(transcribe_listMedicalScribeJobsCmd)
}
