package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_listKnowledgeBasesCmd = &cobra.Command{
	Use:   "list-knowledge-bases",
	Short: "Lists the knowledge bases in an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_listKnowledgeBasesCmd).Standalone()

	bedrockAgent_listKnowledgeBasesCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
	bedrockAgent_listKnowledgeBasesCmd.Flags().String("next-token", "", "If the total number of results is greater than the `maxResults` value provided in the request, enter the token returned in the `nextToken` field in the response in this field to return the next batch of results.")
	bedrockAgentCmd.AddCommand(bedrockAgent_listKnowledgeBasesCmd)
}
