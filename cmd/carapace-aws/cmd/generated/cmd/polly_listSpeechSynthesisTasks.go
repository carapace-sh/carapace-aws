package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var polly_listSpeechSynthesisTasksCmd = &cobra.Command{
	Use:   "list-speech-synthesis-tasks",
	Short: "Returns a list of SpeechSynthesisTask objects ordered by their creation date.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(polly_listSpeechSynthesisTasksCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(polly_listSpeechSynthesisTasksCmd).Standalone()

		polly_listSpeechSynthesisTasksCmd.Flags().String("max-results", "", "Maximum number of speech synthesis tasks returned in a List operation.")
		polly_listSpeechSynthesisTasksCmd.Flags().String("next-token", "", "The pagination token to use in the next request to continue the listing of speech synthesis tasks.")
		polly_listSpeechSynthesisTasksCmd.Flags().String("status", "", "Status of the speech synthesis tasks returned in a List operation")
	})
	pollyCmd.AddCommand(polly_listSpeechSynthesisTasksCmd)
}
