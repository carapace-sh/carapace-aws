package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transcribe_deleteTranscriptionJobCmd = &cobra.Command{
	Use:   "delete-transcription-job",
	Short: "Deletes a transcription job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transcribe_deleteTranscriptionJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(transcribe_deleteTranscriptionJobCmd).Standalone()

		transcribe_deleteTranscriptionJobCmd.Flags().String("transcription-job-name", "", "The name of the transcription job you want to delete.")
		transcribe_deleteTranscriptionJobCmd.MarkFlagRequired("transcription-job-name")
	})
	transcribeCmd.AddCommand(transcribe_deleteTranscriptionJobCmd)
}
