package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var polly_synthesizeSpeechCmd = &cobra.Command{
	Use:   "synthesize-speech",
	Short: "Synthesizes UTF-8 input, plain text or SSML, to a stream of bytes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(polly_synthesizeSpeechCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(polly_synthesizeSpeechCmd).Standalone()

		polly_synthesizeSpeechCmd.Flags().String("engine", "", "Specifies the engine (`standard`, `neural`, `long-form`, or `generative`) for Amazon Polly to use when processing input text for speech synthesis.")
		polly_synthesizeSpeechCmd.Flags().String("language-code", "", "Optional language code for the Synthesize Speech request.")
		polly_synthesizeSpeechCmd.Flags().String("lexicon-names", "", "List of one or more pronunciation lexicon names you want the service to apply during synthesis.")
		polly_synthesizeSpeechCmd.Flags().String("output-format", "", "The format in which the returned output will be encoded.")
		polly_synthesizeSpeechCmd.Flags().String("sample-rate", "", "The audio frequency specified in Hz.")
		polly_synthesizeSpeechCmd.Flags().String("speech-mark-types", "", "The type of speech marks returned for the input text.")
		polly_synthesizeSpeechCmd.Flags().String("text", "", "Input text to synthesize.")
		polly_synthesizeSpeechCmd.Flags().String("text-type", "", "Specifies whether the input text is plain text or SSML.")
		polly_synthesizeSpeechCmd.Flags().String("voice-id", "", "Voice ID to use for the synthesis.")
		polly_synthesizeSpeechCmd.MarkFlagRequired("output-format")
		polly_synthesizeSpeechCmd.MarkFlagRequired("text")
		polly_synthesizeSpeechCmd.MarkFlagRequired("voice-id")
	})
	pollyCmd.AddCommand(polly_synthesizeSpeechCmd)
}
