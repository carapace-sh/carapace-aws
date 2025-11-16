package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcoreControl_getAgentRuntimeEndpointCmd = &cobra.Command{
	Use:   "get-agent-runtime-endpoint",
	Short: "Gets information about an Amazon Secure AgentEndpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcoreControl_getAgentRuntimeEndpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgentcoreControl_getAgentRuntimeEndpointCmd).Standalone()

		bedrockAgentcoreControl_getAgentRuntimeEndpointCmd.Flags().String("agent-runtime-id", "", "The unique identifier of the AgentCore Runtime associated with the endpoint.")
		bedrockAgentcoreControl_getAgentRuntimeEndpointCmd.Flags().String("endpoint-name", "", "The name of the AgentCore Runtime endpoint to retrieve.")
		bedrockAgentcoreControl_getAgentRuntimeEndpointCmd.MarkFlagRequired("agent-runtime-id")
		bedrockAgentcoreControl_getAgentRuntimeEndpointCmd.MarkFlagRequired("endpoint-name")
	})
	bedrockAgentcoreControlCmd.AddCommand(bedrockAgentcoreControl_getAgentRuntimeEndpointCmd)
}
