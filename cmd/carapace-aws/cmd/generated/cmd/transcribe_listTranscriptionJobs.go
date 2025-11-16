package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transcribe_listTranscriptionJobsCmd = &cobra.Command{
	Use:   "list-transcription-jobs",
	Short: "Provides a list of transcription jobs that match the specified criteria.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transcribe_listTranscriptionJobsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(transcribe_listTranscriptionJobsCmd).Standalone()

		transcribe_listTranscriptionJobsCmd.Flags().String("job-name-contains", "", "Returns only the transcription jobs that contain the specified string.")
		transcribe_listTranscriptionJobsCmd.Flags().String("max-results", "", "The maximum number of transcription jobs to return in each page of results.")
		transcribe_listTranscriptionJobsCmd.Flags().String("next-token", "", "If your `ListTranscriptionJobs` request returns more results than can be displayed, `NextToken` is displayed in the response with an associated string.")
		transcribe_listTranscriptionJobsCmd.Flags().String("status", "", "Returns only transcription jobs with the specified status.")
	})
	transcribeCmd.AddCommand(transcribe_listTranscriptionJobsCmd)
}
