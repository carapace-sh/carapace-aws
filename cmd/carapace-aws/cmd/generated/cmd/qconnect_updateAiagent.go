package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_updateAiagentCmd = &cobra.Command{
	Use:   "update-aiagent",
	Short: "Updates an AI Agent.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_updateAiagentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qconnect_updateAiagentCmd).Standalone()

		qconnect_updateAiagentCmd.Flags().String("ai-agent-id", "", "The identifier of the Amazon Q in Connect AI Agent.")
		qconnect_updateAiagentCmd.Flags().String("assistant-id", "", "The identifier of the Amazon Q in Connect assistant.")
		qconnect_updateAiagentCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		qconnect_updateAiagentCmd.Flags().String("configuration", "", "The configuration of the Amazon Q in Connect AI Agent.")
		qconnect_updateAiagentCmd.Flags().String("description", "", "The description of the Amazon Q in Connect AI Agent.")
		qconnect_updateAiagentCmd.Flags().String("visibility-status", "", "The visbility status of the Amazon Q in Connect AI Agent.")
		qconnect_updateAiagentCmd.MarkFlagRequired("ai-agent-id")
		qconnect_updateAiagentCmd.MarkFlagRequired("assistant-id")
		qconnect_updateAiagentCmd.MarkFlagRequired("visibility-status")
	})
	qconnectCmd.AddCommand(qconnect_updateAiagentCmd)
}
