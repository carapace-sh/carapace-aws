package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transcribe_startMedicalTranscriptionJobCmd = &cobra.Command{
	Use:   "start-medical-transcription-job",
	Short: "Transcribes the audio from a medical dictation or conversation and applies any additional Request Parameters you choose to include in your request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transcribe_startMedicalTranscriptionJobCmd).Standalone()

	transcribe_startMedicalTranscriptionJobCmd.Flags().String("content-identification-type", "", "Labels all personal health information (PHI) identified in your transcript.")
	transcribe_startMedicalTranscriptionJobCmd.Flags().String("kmsencryption-context", "", "A map of plain text, non-secret key:value pairs, known as encryption context pairs, that provide an added layer of security for your data.")
	transcribe_startMedicalTranscriptionJobCmd.Flags().String("language-code", "", "The language code that represents the language spoken in the input media file.")
	transcribe_startMedicalTranscriptionJobCmd.Flags().String("media", "", "")
	transcribe_startMedicalTranscriptionJobCmd.Flags().String("media-format", "", "Specify the format of your input media file.")
	transcribe_startMedicalTranscriptionJobCmd.Flags().String("media-sample-rate-hertz", "", "The sample rate, in hertz, of the audio track in your input media file.")
	transcribe_startMedicalTranscriptionJobCmd.Flags().String("medical-transcription-job-name", "", "A unique name, chosen by you, for your medical transcription job.")
	transcribe_startMedicalTranscriptionJobCmd.Flags().String("output-bucket-name", "", "The name of the Amazon S3 bucket where you want your medical transcription output stored.")
	transcribe_startMedicalTranscriptionJobCmd.Flags().String("output-encryption-kmskey-id", "", "The Amazon Resource Name (ARN) of a KMS key that you want to use to encrypt your medical transcription output.")
	transcribe_startMedicalTranscriptionJobCmd.Flags().String("output-key", "", "Use in combination with `OutputBucketName` to specify the output location of your transcript and, optionally, a unique name for your output file.")
	transcribe_startMedicalTranscriptionJobCmd.Flags().String("settings", "", "Specify additional optional settings in your request, including channel identification, alternative transcriptions, and speaker partitioning.")
	transcribe_startMedicalTranscriptionJobCmd.Flags().String("specialty", "", "Specify the predominant medical specialty represented in your media.")
	transcribe_startMedicalTranscriptionJobCmd.Flags().String("tags", "", "Adds one or more custom tags, each in the form of a key:value pair, to a new medical transcription job at the time you start this new job.")
	transcribe_startMedicalTranscriptionJobCmd.Flags().String("type", "", "Specify whether your input media contains only one person (`DICTATION`) or contains a conversation between two people (`CONVERSATION`).")
	transcribe_startMedicalTranscriptionJobCmd.MarkFlagRequired("language-code")
	transcribe_startMedicalTranscriptionJobCmd.MarkFlagRequired("media")
	transcribe_startMedicalTranscriptionJobCmd.MarkFlagRequired("medical-transcription-job-name")
	transcribe_startMedicalTranscriptionJobCmd.MarkFlagRequired("output-bucket-name")
	transcribe_startMedicalTranscriptionJobCmd.MarkFlagRequired("specialty")
	transcribe_startMedicalTranscriptionJobCmd.MarkFlagRequired("type")
	transcribeCmd.AddCommand(transcribe_startMedicalTranscriptionJobCmd)
}
