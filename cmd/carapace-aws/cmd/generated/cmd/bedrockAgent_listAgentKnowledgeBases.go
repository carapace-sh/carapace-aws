package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_listAgentKnowledgeBasesCmd = &cobra.Command{
	Use:   "list-agent-knowledge-bases",
	Short: "Lists knowledge bases associated with an agent and information about each one.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_listAgentKnowledgeBasesCmd).Standalone()

	bedrockAgent_listAgentKnowledgeBasesCmd.Flags().String("agent-id", "", "The unique identifier of the agent for which to return information about knowledge bases associated with it.")
	bedrockAgent_listAgentKnowledgeBasesCmd.Flags().String("agent-version", "", "The version of the agent for which to return information about knowledge bases associated with it.")
	bedrockAgent_listAgentKnowledgeBasesCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
	bedrockAgent_listAgentKnowledgeBasesCmd.Flags().String("next-token", "", "If the total number of results is greater than the `maxResults` value provided in the request, enter the token returned in the `nextToken` field in the response in this field to return the next batch of results.")
	bedrockAgent_listAgentKnowledgeBasesCmd.MarkFlagRequired("agent-id")
	bedrockAgent_listAgentKnowledgeBasesCmd.MarkFlagRequired("agent-version")
	bedrockAgentCmd.AddCommand(bedrockAgent_listAgentKnowledgeBasesCmd)
}
