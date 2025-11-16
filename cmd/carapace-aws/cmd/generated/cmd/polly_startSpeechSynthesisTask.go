package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var polly_startSpeechSynthesisTaskCmd = &cobra.Command{
	Use:   "start-speech-synthesis-task",
	Short: "Allows the creation of an asynchronous synthesis task, by starting a new `SpeechSynthesisTask`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(polly_startSpeechSynthesisTaskCmd).Standalone()

	polly_startSpeechSynthesisTaskCmd.Flags().String("engine", "", "Specifies the engine (`standard`, `neural`, `long-form` or `generative`) for Amazon Polly to use when processing input text for speech synthesis.")
	polly_startSpeechSynthesisTaskCmd.Flags().String("language-code", "", "Optional language code for the Speech Synthesis request.")
	polly_startSpeechSynthesisTaskCmd.Flags().String("lexicon-names", "", "List of one or more pronunciation lexicon names you want the service to apply during synthesis.")
	polly_startSpeechSynthesisTaskCmd.Flags().String("output-format", "", "The format in which the returned output will be encoded.")
	polly_startSpeechSynthesisTaskCmd.Flags().String("output-s3-bucket-name", "", "Amazon S3 bucket name to which the output file will be saved.")
	polly_startSpeechSynthesisTaskCmd.Flags().String("output-s3-key-prefix", "", "The Amazon S3 key prefix for the output speech file.")
	polly_startSpeechSynthesisTaskCmd.Flags().String("sample-rate", "", "The audio frequency specified in Hz.")
	polly_startSpeechSynthesisTaskCmd.Flags().String("sns-topic-arn", "", "ARN for the SNS topic optionally used for providing status notification for a speech synthesis task.")
	polly_startSpeechSynthesisTaskCmd.Flags().String("speech-mark-types", "", "The type of speech marks returned for the input text.")
	polly_startSpeechSynthesisTaskCmd.Flags().String("text", "", "The input text to synthesize.")
	polly_startSpeechSynthesisTaskCmd.Flags().String("text-type", "", "Specifies whether the input text is plain text or SSML.")
	polly_startSpeechSynthesisTaskCmd.Flags().String("voice-id", "", "Voice ID to use for the synthesis.")
	polly_startSpeechSynthesisTaskCmd.MarkFlagRequired("output-format")
	polly_startSpeechSynthesisTaskCmd.MarkFlagRequired("output-s3-bucket-name")
	polly_startSpeechSynthesisTaskCmd.MarkFlagRequired("text")
	polly_startSpeechSynthesisTaskCmd.MarkFlagRequired("voice-id")
	pollyCmd.AddCommand(polly_startSpeechSynthesisTaskCmd)
}
