package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_listAiagentsCmd = &cobra.Command{
	Use:   "list-aiagents",
	Short: "Lists AI Agents.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_listAiagentsCmd).Standalone()

	qconnect_listAiagentsCmd.Flags().String("assistant-id", "", "The identifier of the Amazon Q in Connect assistant.")
	qconnect_listAiagentsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	qconnect_listAiagentsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	qconnect_listAiagentsCmd.Flags().String("origin", "", "The origin of the AI Agents to be listed.")
	qconnect_listAiagentsCmd.MarkFlagRequired("assistant-id")
	qconnectCmd.AddCommand(qconnect_listAiagentsCmd)
}
