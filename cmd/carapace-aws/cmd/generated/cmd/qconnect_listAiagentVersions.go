package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_listAiagentVersionsCmd = &cobra.Command{
	Use:   "list-aiagent-versions",
	Short: "List AI Agent versions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_listAiagentVersionsCmd).Standalone()

	qconnect_listAiagentVersionsCmd.Flags().String("ai-agent-id", "", "The identifier of the Amazon Q in Connect AI Agent for which versions are to be listed.")
	qconnect_listAiagentVersionsCmd.Flags().String("assistant-id", "", "The identifier of the Amazon Q in Connect assistant.")
	qconnect_listAiagentVersionsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	qconnect_listAiagentVersionsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	qconnect_listAiagentVersionsCmd.Flags().String("origin", "", "The origin of the AI Agent versions to be listed.")
	qconnect_listAiagentVersionsCmd.MarkFlagRequired("ai-agent-id")
	qconnect_listAiagentVersionsCmd.MarkFlagRequired("assistant-id")
	qconnectCmd.AddCommand(qconnect_listAiagentVersionsCmd)
}
