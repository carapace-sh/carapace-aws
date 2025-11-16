package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transcribe_getTranscriptionJobCmd = &cobra.Command{
	Use:   "get-transcription-job",
	Short: "Provides information about the specified transcription job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transcribe_getTranscriptionJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(transcribe_getTranscriptionJobCmd).Standalone()

		transcribe_getTranscriptionJobCmd.Flags().String("transcription-job-name", "", "The name of the transcription job you want information about.")
		transcribe_getTranscriptionJobCmd.MarkFlagRequired("transcription-job-name")
	})
	transcribeCmd.AddCommand(transcribe_getTranscriptionJobCmd)
}
