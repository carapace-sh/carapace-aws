package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_listKnowledgeBasesCmd = &cobra.Command{
	Use:   "list-knowledge-bases",
	Short: "Lists the knowledge bases.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_listKnowledgeBasesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qconnect_listKnowledgeBasesCmd).Standalone()

		qconnect_listKnowledgeBasesCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		qconnect_listKnowledgeBasesCmd.Flags().String("next-token", "", "The token for the next set of results.")
	})
	qconnectCmd.AddCommand(qconnect_listKnowledgeBasesCmd)
}
