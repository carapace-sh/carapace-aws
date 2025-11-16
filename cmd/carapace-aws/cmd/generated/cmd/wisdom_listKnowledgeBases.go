package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wisdom_listKnowledgeBasesCmd = &cobra.Command{
	Use:   "list-knowledge-bases",
	Short: "Lists the knowledge bases.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wisdom_listKnowledgeBasesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wisdom_listKnowledgeBasesCmd).Standalone()

		wisdom_listKnowledgeBasesCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		wisdom_listKnowledgeBasesCmd.Flags().String("next-token", "", "The token for the next set of results.")
	})
	wisdomCmd.AddCommand(wisdom_listKnowledgeBasesCmd)
}
