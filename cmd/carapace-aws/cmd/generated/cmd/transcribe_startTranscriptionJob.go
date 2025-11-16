package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transcribe_startTranscriptionJobCmd = &cobra.Command{
	Use:   "start-transcription-job",
	Short: "Transcribes the audio from a media file and applies any additional Request Parameters you choose to include in your request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transcribe_startTranscriptionJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(transcribe_startTranscriptionJobCmd).Standalone()

		transcribe_startTranscriptionJobCmd.Flags().String("content-redaction", "", "Makes it possible to redact or flag specified personally identifiable information (PII) in your transcript.")
		transcribe_startTranscriptionJobCmd.Flags().Bool("identify-language", false, "Enables automatic language identification in your transcription job request.")
		transcribe_startTranscriptionJobCmd.Flags().Bool("identify-multiple-languages", false, "Enables automatic multi-language identification in your transcription job request.")
		transcribe_startTranscriptionJobCmd.Flags().String("job-execution-settings", "", "Makes it possible to control how your transcription job is processed.")
		transcribe_startTranscriptionJobCmd.Flags().String("kmsencryption-context", "", "A map of plain text, non-secret key:value pairs, known as encryption context pairs, that provide an added layer of security for your data.")
		transcribe_startTranscriptionJobCmd.Flags().String("language-code", "", "The language code that represents the language spoken in the input media file.")
		transcribe_startTranscriptionJobCmd.Flags().String("language-id-settings", "", "If using automatic language identification in your request and you want to apply a custom language model, a custom vocabulary, or a custom vocabulary filter, include `LanguageIdSettings` with the relevant sub-parameters (`VocabularyName`, `LanguageModelName`, and `VocabularyFilterName`).")
		transcribe_startTranscriptionJobCmd.Flags().String("language-options", "", "You can specify two or more language codes that represent the languages you think may be present in your media.")
		transcribe_startTranscriptionJobCmd.Flags().String("media", "", "Describes the Amazon S3 location of the media file you want to use in your request.")
		transcribe_startTranscriptionJobCmd.Flags().String("media-format", "", "Specify the format of your input media file.")
		transcribe_startTranscriptionJobCmd.Flags().String("media-sample-rate-hertz", "", "The sample rate, in hertz, of the audio track in your input media file.")
		transcribe_startTranscriptionJobCmd.Flags().String("model-settings", "", "Specify the custom language model you want to include with your transcription job.")
		transcribe_startTranscriptionJobCmd.Flags().Bool("no-identify-language", false, "Enables automatic language identification in your transcription job request.")
		transcribe_startTranscriptionJobCmd.Flags().Bool("no-identify-multiple-languages", false, "Enables automatic multi-language identification in your transcription job request.")
		transcribe_startTranscriptionJobCmd.Flags().String("output-bucket-name", "", "The name of the Amazon S3 bucket where you want your transcription output stored.")
		transcribe_startTranscriptionJobCmd.Flags().String("output-encryption-kmskey-id", "", "The Amazon Resource Name (ARN) of a KMS key that you want to use to encrypt your transcription output.")
		transcribe_startTranscriptionJobCmd.Flags().String("output-key", "", "Use in combination with `OutputBucketName` to specify the output location of your transcript and, optionally, a unique name for your output file.")
		transcribe_startTranscriptionJobCmd.Flags().String("settings", "", "Specify additional optional settings in your request, including channel identification, alternative transcriptions, speaker partitioning.")
		transcribe_startTranscriptionJobCmd.Flags().String("subtitles", "", "Produces subtitle files for your input media.")
		transcribe_startTranscriptionJobCmd.Flags().String("tags", "", "Adds one or more custom tags, each in the form of a key:value pair, to a new transcription job at the time you start this new job.")
		transcribe_startTranscriptionJobCmd.Flags().String("toxicity-detection", "", "Enables toxic speech detection in your transcript.")
		transcribe_startTranscriptionJobCmd.Flags().String("transcription-job-name", "", "A unique name, chosen by you, for your transcription job.")
		transcribe_startTranscriptionJobCmd.MarkFlagRequired("media")
		transcribe_startTranscriptionJobCmd.Flag("no-identify-language").Hidden = true
		transcribe_startTranscriptionJobCmd.Flag("no-identify-multiple-languages").Hidden = true
		transcribe_startTranscriptionJobCmd.MarkFlagRequired("transcription-job-name")
	})
	transcribeCmd.AddCommand(transcribe_startTranscriptionJobCmd)
}
