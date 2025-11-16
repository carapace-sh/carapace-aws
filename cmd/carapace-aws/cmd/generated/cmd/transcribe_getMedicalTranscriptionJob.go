package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transcribe_getMedicalTranscriptionJobCmd = &cobra.Command{
	Use:   "get-medical-transcription-job",
	Short: "Provides information about the specified medical transcription job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transcribe_getMedicalTranscriptionJobCmd).Standalone()

	transcribe_getMedicalTranscriptionJobCmd.Flags().String("medical-transcription-job-name", "", "The name of the medical transcription job you want information about.")
	transcribe_getMedicalTranscriptionJobCmd.MarkFlagRequired("medical-transcription-job-name")
	transcribeCmd.AddCommand(transcribe_getMedicalTranscriptionJobCmd)
}
