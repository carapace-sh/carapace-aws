package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_removeAssistantAiagentCmd = &cobra.Command{
	Use:   "remove-assistant-aiagent",
	Short: "Removes the AI Agent that is set for use by default on an Amazon Q in Connect Assistant.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_removeAssistantAiagentCmd).Standalone()

	qconnect_removeAssistantAiagentCmd.Flags().String("ai-agent-type", "", "The type of the AI Agent being removed for use by default from the Amazon Q in Connect Assistant.")
	qconnect_removeAssistantAiagentCmd.Flags().String("assistant-id", "", "The identifier of the Amazon Q in Connect assistant.")
	qconnect_removeAssistantAiagentCmd.MarkFlagRequired("ai-agent-type")
	qconnect_removeAssistantAiagentCmd.MarkFlagRequired("assistant-id")
	qconnectCmd.AddCommand(qconnect_removeAssistantAiagentCmd)
}
