package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_updateAssistantAiagentCmd = &cobra.Command{
	Use:   "update-assistant-aiagent",
	Short: "Updates the AI Agent that is set for use by default on an Amazon Q in Connect Assistant.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_updateAssistantAiagentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qconnect_updateAssistantAiagentCmd).Standalone()

		qconnect_updateAssistantAiagentCmd.Flags().String("ai-agent-type", "", "The type of the AI Agent being updated for use by default on the Amazon Q in Connect Assistant.")
		qconnect_updateAssistantAiagentCmd.Flags().String("assistant-id", "", "The identifier of the Amazon Q in Connect assistant.")
		qconnect_updateAssistantAiagentCmd.Flags().String("configuration", "", "The configuration of the AI Agent being updated for use by default on the Amazon Q in Connect Assistant.")
		qconnect_updateAssistantAiagentCmd.MarkFlagRequired("ai-agent-type")
		qconnect_updateAssistantAiagentCmd.MarkFlagRequired("assistant-id")
		qconnect_updateAssistantAiagentCmd.MarkFlagRequired("configuration")
	})
	qconnectCmd.AddCommand(qconnect_updateAssistantAiagentCmd)
}
