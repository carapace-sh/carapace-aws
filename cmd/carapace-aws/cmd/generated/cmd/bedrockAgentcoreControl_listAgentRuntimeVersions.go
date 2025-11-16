package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcoreControl_listAgentRuntimeVersionsCmd = &cobra.Command{
	Use:   "list-agent-runtime-versions",
	Short: "Lists all versions of a specific Amazon Secure Agent.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcoreControl_listAgentRuntimeVersionsCmd).Standalone()

	bedrockAgentcoreControl_listAgentRuntimeVersionsCmd.Flags().String("agent-runtime-id", "", "The unique identifier of the AgentCore Runtime to list versions for.")
	bedrockAgentcoreControl_listAgentRuntimeVersionsCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
	bedrockAgentcoreControl_listAgentRuntimeVersionsCmd.Flags().String("next-token", "", "A token to retrieve the next page of results.")
	bedrockAgentcoreControl_listAgentRuntimeVersionsCmd.MarkFlagRequired("agent-runtime-id")
	bedrockAgentcoreControlCmd.AddCommand(bedrockAgentcoreControl_listAgentRuntimeVersionsCmd)
}
