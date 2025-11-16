package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_deleteAiagentVersionCmd = &cobra.Command{
	Use:   "delete-aiagent-version",
	Short: "Deletes an Amazon Q in Connect AI Agent Version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_deleteAiagentVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qconnect_deleteAiagentVersionCmd).Standalone()

		qconnect_deleteAiagentVersionCmd.Flags().String("ai-agent-id", "", "The identifier of the Amazon Q in Connect AI Agent.")
		qconnect_deleteAiagentVersionCmd.Flags().String("assistant-id", "", "The identifier of the Amazon Q in Connect assistant.")
		qconnect_deleteAiagentVersionCmd.Flags().String("version-number", "", "The version number of the AI Agent version.")
		qconnect_deleteAiagentVersionCmd.MarkFlagRequired("ai-agent-id")
		qconnect_deleteAiagentVersionCmd.MarkFlagRequired("assistant-id")
		qconnect_deleteAiagentVersionCmd.MarkFlagRequired("version-number")
	})
	qconnectCmd.AddCommand(qconnect_deleteAiagentVersionCmd)
}
