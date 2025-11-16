package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_listAgentCollaboratorsCmd = &cobra.Command{
	Use:   "list-agent-collaborators",
	Short: "Retrieve a list of an agent's collaborators.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_listAgentCollaboratorsCmd).Standalone()

	bedrockAgent_listAgentCollaboratorsCmd.Flags().String("agent-id", "", "The agent's ID.")
	bedrockAgent_listAgentCollaboratorsCmd.Flags().String("agent-version", "", "The agent's version.")
	bedrockAgent_listAgentCollaboratorsCmd.Flags().String("max-results", "", "The maximum number of agent collaborators to return in one page of results.")
	bedrockAgent_listAgentCollaboratorsCmd.Flags().String("next-token", "", "Specify the pagination token from a previous request to retrieve the next page of results.")
	bedrockAgent_listAgentCollaboratorsCmd.MarkFlagRequired("agent-id")
	bedrockAgent_listAgentCollaboratorsCmd.MarkFlagRequired("agent-version")
	bedrockAgentCmd.AddCommand(bedrockAgent_listAgentCollaboratorsCmd)
}
