package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_createAiagentVersionCmd = &cobra.Command{
	Use:   "create-aiagent-version",
	Short: "Creates and Amazon Q in Connect AI Agent version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_createAiagentVersionCmd).Standalone()

	qconnect_createAiagentVersionCmd.Flags().String("ai-agent-id", "", "The identifier of the Amazon Q in Connect AI Agent.")
	qconnect_createAiagentVersionCmd.Flags().String("assistant-id", "", "The identifier of the Amazon Q in Connect assistant.")
	qconnect_createAiagentVersionCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	qconnect_createAiagentVersionCmd.Flags().String("modified-time", "", "The modification time of the AI Agent should be tracked for version creation.")
	qconnect_createAiagentVersionCmd.MarkFlagRequired("ai-agent-id")
	qconnect_createAiagentVersionCmd.MarkFlagRequired("assistant-id")
	qconnectCmd.AddCommand(qconnect_createAiagentVersionCmd)
}
