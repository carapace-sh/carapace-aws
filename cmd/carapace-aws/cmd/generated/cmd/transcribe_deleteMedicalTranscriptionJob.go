package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transcribe_deleteMedicalTranscriptionJobCmd = &cobra.Command{
	Use:   "delete-medical-transcription-job",
	Short: "Deletes a medical transcription job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transcribe_deleteMedicalTranscriptionJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(transcribe_deleteMedicalTranscriptionJobCmd).Standalone()

		transcribe_deleteMedicalTranscriptionJobCmd.Flags().String("medical-transcription-job-name", "", "The name of the medical transcription job you want to delete.")
		transcribe_deleteMedicalTranscriptionJobCmd.MarkFlagRequired("medical-transcription-job-name")
	})
	transcribeCmd.AddCommand(transcribe_deleteMedicalTranscriptionJobCmd)
}
