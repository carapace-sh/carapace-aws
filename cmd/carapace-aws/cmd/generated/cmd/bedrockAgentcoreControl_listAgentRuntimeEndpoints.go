package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcoreControl_listAgentRuntimeEndpointsCmd = &cobra.Command{
	Use:   "list-agent-runtime-endpoints",
	Short: "Lists all endpoints for a specific Amazon Secure Agent.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcoreControl_listAgentRuntimeEndpointsCmd).Standalone()

	bedrockAgentcoreControl_listAgentRuntimeEndpointsCmd.Flags().String("agent-runtime-id", "", "The unique identifier of the AgentCore Runtime to list endpoints for.")
	bedrockAgentcoreControl_listAgentRuntimeEndpointsCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
	bedrockAgentcoreControl_listAgentRuntimeEndpointsCmd.Flags().String("next-token", "", "A token to retrieve the next page of results.")
	bedrockAgentcoreControl_listAgentRuntimeEndpointsCmd.MarkFlagRequired("agent-runtime-id")
	bedrockAgentcoreControlCmd.AddCommand(bedrockAgentcoreControl_listAgentRuntimeEndpointsCmd)
}
