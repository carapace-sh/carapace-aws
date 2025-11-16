package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_getAiagentCmd = &cobra.Command{
	Use:   "get-aiagent",
	Short: "Gets an Amazon Q in Connect AI Agent.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_getAiagentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qconnect_getAiagentCmd).Standalone()

		qconnect_getAiagentCmd.Flags().String("ai-agent-id", "", "The identifier of the Amazon Q in Connect AI Agent (with or without a version qualifier).")
		qconnect_getAiagentCmd.Flags().String("assistant-id", "", "The identifier of the Amazon Q in Connect assistant.")
		qconnect_getAiagentCmd.MarkFlagRequired("ai-agent-id")
		qconnect_getAiagentCmd.MarkFlagRequired("assistant-id")
	})
	qconnectCmd.AddCommand(qconnect_getAiagentCmd)
}
