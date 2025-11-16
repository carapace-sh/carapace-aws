package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transcribe_listMedicalTranscriptionJobsCmd = &cobra.Command{
	Use:   "list-medical-transcription-jobs",
	Short: "Provides a list of medical transcription jobs that match the specified criteria.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transcribe_listMedicalTranscriptionJobsCmd).Standalone()

	transcribe_listMedicalTranscriptionJobsCmd.Flags().String("job-name-contains", "", "Returns only the medical transcription jobs that contain the specified string.")
	transcribe_listMedicalTranscriptionJobsCmd.Flags().String("max-results", "", "The maximum number of medical transcription jobs to return in each page of results.")
	transcribe_listMedicalTranscriptionJobsCmd.Flags().String("next-token", "", "If your `ListMedicalTranscriptionJobs` request returns more results than can be displayed, `NextToken` is displayed in the response with an associated string.")
	transcribe_listMedicalTranscriptionJobsCmd.Flags().String("status", "", "Returns only medical transcription jobs with the specified status.")
	transcribeCmd.AddCommand(transcribe_listMedicalTranscriptionJobsCmd)
}
