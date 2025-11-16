package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var voiceId_evaluateSessionCmd = &cobra.Command{
	Use:   "evaluate-session",
	Short: "Evaluates a specified session based on audio data accumulated during a streaming Amazon Connect Voice ID call.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(voiceId_evaluateSessionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(voiceId_evaluateSessionCmd).Standalone()

		voiceId_evaluateSessionCmd.Flags().String("domain-id", "", "The identifier of the domain where the session started.")
		voiceId_evaluateSessionCmd.Flags().String("session-name-or-id", "", "The session identifier, or name of the session, that you want to evaluate.")
		voiceId_evaluateSessionCmd.MarkFlagRequired("domain-id")
		voiceId_evaluateSessionCmd.MarkFlagRequired("session-name-or-id")
	})
	voiceIdCmd.AddCommand(voiceId_evaluateSessionCmd)
}
