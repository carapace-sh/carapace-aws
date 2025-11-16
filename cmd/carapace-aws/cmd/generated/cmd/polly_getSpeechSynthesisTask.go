package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var polly_getSpeechSynthesisTaskCmd = &cobra.Command{
	Use:   "get-speech-synthesis-task",
	Short: "Retrieves a specific SpeechSynthesisTask object based on its TaskID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(polly_getSpeechSynthesisTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(polly_getSpeechSynthesisTaskCmd).Standalone()

		polly_getSpeechSynthesisTaskCmd.Flags().String("task-id", "", "The Amazon Polly generated identifier for a speech synthesis task.")
		polly_getSpeechSynthesisTaskCmd.MarkFlagRequired("task-id")
	})
	pollyCmd.AddCommand(polly_getSpeechSynthesisTaskCmd)
}
