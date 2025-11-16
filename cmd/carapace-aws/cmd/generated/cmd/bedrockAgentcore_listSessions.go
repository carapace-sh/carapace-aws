package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcore_listSessionsCmd = &cobra.Command{
	Use:   "list-sessions",
	Short: "Lists sessions in an AgentCore Memory resource based on specified criteria.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcore_listSessionsCmd).Standalone()

	bedrockAgentcore_listSessionsCmd.Flags().String("actor-id", "", "The identifier of the actor for which to list sessions.")
	bedrockAgentcore_listSessionsCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
	bedrockAgentcore_listSessionsCmd.Flags().String("memory-id", "", "The identifier of the AgentCore Memory resource for which to list sessions.")
	bedrockAgentcore_listSessionsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	bedrockAgentcore_listSessionsCmd.MarkFlagRequired("actor-id")
	bedrockAgentcore_listSessionsCmd.MarkFlagRequired("memory-id")
	bedrockAgentcoreCmd.AddCommand(bedrockAgentcore_listSessionsCmd)
}
