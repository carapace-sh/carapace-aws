package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_deleteAiagentCmd = &cobra.Command{
	Use:   "delete-aiagent",
	Short: "Deletes an Amazon Q in Connect AI Agent.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_deleteAiagentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qconnect_deleteAiagentCmd).Standalone()

		qconnect_deleteAiagentCmd.Flags().String("ai-agent-id", "", "The identifier of the Amazon Q in Connect AI Agent.")
		qconnect_deleteAiagentCmd.Flags().String("assistant-id", "", "The identifier of the Amazon Q in Connect assistant.")
		qconnect_deleteAiagentCmd.MarkFlagRequired("ai-agent-id")
		qconnect_deleteAiagentCmd.MarkFlagRequired("assistant-id")
	})
	qconnectCmd.AddCommand(qconnect_deleteAiagentCmd)
}
